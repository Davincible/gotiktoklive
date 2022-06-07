package gotiktoklive

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	pb "github.com/Davincible/gotiktoklive/proto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"google.golang.org/protobuf/proto"
)

// TODO: check gift prices of gifts not in wish list

const (
	POLLING_INTERVAL         = time.Second
	DEFAULT_EVENTS_CHAN_SIZE = 100
)

// Live allows you to track a livestream.
// To track a user call tiktok.TrackUser(<user>).
type Live struct {
	t *TikTok

	cursor   string
	wss      net.Conn
	wsURL    string
	wsParams map[string]string
	close    func()
	done     func() <-chan struct{}

	ID       string
	Info     *RoomInfo
	GiftInfo *GiftInfo
	Events   chan interface{}
	chanSize int
}

func (t *TikTok) newLive(roomId string) *Live {
	live := Live{
		t:        t,
		ID:       roomId,
		Events:   make(chan interface{}, DEFAULT_EVENTS_CHAN_SIZE),
		chanSize: DEFAULT_EVENTS_CHAN_SIZE,
	}

	ctx, cancel := context.WithCancel(context.Background())
	live.done = ctx.Done
	o := sync.Once{}
	live.close = func() {
		o.Do(func() {
			cancel()
			t.wg.Wait()
			close(live.Events)
		})
	}

	return &live
}

// Close will terminate the connection and stop any downloads.
func (l *Live) Close() {
	l.close()
}

func (l *Live) fetchRoom() error {
	roomInfo, err := l.getRoomInfo()
	if err != nil {
		return err
	}
	l.Info = roomInfo

	giftInfo, err := l.getGiftInfo()
	if err != nil {
		return err
	}
	l.GiftInfo = giftInfo

	err = l.getRoomData()
	if err != nil {
		return err
	}
	return nil
}

// GetRoomInfo will only fetch the room info, normally available with live.Info
//  but not start tracking a live stream.
func (t *TikTok) GetRoomInfo(username string) (*RoomInfo, error) {
	id, err := t.getRoomID(username)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch room ID by username")
	}

	l := Live{
		t:  t,
		ID: id,
	}

	roomInfo, err := l.getRoomInfo()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch room info")
	}
	return roomInfo, nil
}

// TrackUser will start to track the livestream of a user, if live.
// To listen to events emitted by the livestream, such as comments and viewer
//  count, listen to the Live.Events channel.
// It will start a go routine and connect to the tiktok websocket.
func (t *TikTok) TrackUser(username string) (*Live, error) {
	id, err := t.getRoomID(username)
	if err != nil {
		return nil, err
	}

	return t.TrackRoom(id)
}

// TrackRoom will start to track a room by room ID.
// It will start a go routine and connect to the tiktok websocket.
func (t *TikTok) TrackRoom(roomId string) (*Live, error) {
	live := t.newLive(roomId)

	if err := live.fetchRoom(); err != nil {
		return nil, err
	}

	if err := live.connectRoom(); err != nil {
		return nil, err
	}

	return live, nil
}

func (live *Live) connectRoom() error {
	wss, err := live.tryConnectionUpgrade()
	if err != nil {
		return err
	}
	if !wss {
		live.t.wg.Add(1)
		live.startPolling()
	}

	return nil
}

func (t *TikTok) getRoomID(user string) (string, error) {
	userInfo, err := t.GetUserInfo(user)
	if err != nil {
		return "", err
	}

	if userInfo.RoomID == "" {
		return "", ErrUserOffline
	}
	return userInfo.RoomID, nil
}

func (l *Live) getRoomInfo() (*RoomInfo, error) {
	t := l.t

	params := copyMap(defaultGETParams)
	params["room_id"] = l.ID

	body, err := t.sendRequest(&reqOptions{
		Endpoint: urlRoomInfo,
		Query:    params,
	})
	if err != nil {
		return nil, err
	}

	var rsp roomInfoRsp
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}

	if rsp.RoomInfo.Status == 4 {
		return rsp.RoomInfo, ErrLiveHasEnded
	}
	return rsp.RoomInfo, nil
}

func (l *Live) getGiftInfo() (*GiftInfo, error) {
	t := l.t

	params := copyMap(defaultGETParams)
	params["room_id"] = l.ID

	body, err := t.sendRequest(&reqOptions{
		Endpoint: urlGiftInfo,
		Query:    params,
	})
	if err != nil {
		return nil, err
	}

	var rsp giftInfoRsp
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return rsp.GiftInfo, nil
}

func (l *Live) getRoomData() error {
	t := l.t

	params := copyMap(defaultGETParams)
	params["room_id"] = l.ID

	if l.cursor != "" {
		params["cursor"] = l.cursor
	}

	body, err := t.sendRequest(&reqOptions{
		Endpoint: urlRoomData,
		Query:    params,
	})
	if err != nil {
		return err
	}
	var rsp pb.WebcastResponse
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return err
	}

	l.cursor = rsp.Cursor
	if rsp.WsUrl != "" && rsp.WsParam != nil {
		l.wsURL = rsp.WsUrl
		l.wsParams = map[string]string{rsp.WsParam.Name: rsp.WsParam.Value}
	}

	for _, msg := range rsp.Messages {
		parsed, err := parseMsg(msg, t.warnHandler)
		if err != nil {
			return err
		}
		l.Events <- parsed
	}

	return nil
}

func (l *Live) startPolling() {
	ticker := time.NewTicker(POLLING_INTERVAL)
	defer ticker.Stop()
	defer l.t.wg.Done()

	var lastUpgradeAttempt time.Time

	l.t.infoHandler("Started polling")

	for {
		select {
		case <-ticker.C:
			err := l.getRoomData()
			if err != nil {
				l.t.errHandler(err)
			}

			if lastUpgradeAttempt.IsZero() || time.Now().Add(-time.Minute*5).Unix() > lastUpgradeAttempt.Unix() {
				lastUpgradeAttempt = time.Now()
				wss, err := l.tryConnectionUpgrade()
				if err != nil {
					l.t.errHandler(err)
				}
				if wss {
					return
				}
			}
		case <-l.t.done():
			l.t.infoHandler("Stopped polling")
			return
		}
	}
}

// DownloadStream will download the stream to an .mkv file.
//
// A filename can be optionally provided as an argument, if not provided one
//  will be generated, with the stream start time in the format of 2022y05m25dT13h03m16s.
// The stream start time can be found in Live.Info.CreateTime as epoch seconds.
func (l *Live) DownloadStream(file ...string) error {
	// Check if ffmpeg is installed
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return ErrFFMPEGNotFound
	}

	// Get URl
	url := l.Info.StreamURL.HlsPullURL
	if url == "" {
		return ErrUrlNotFound
	}

	// Set file path
	var path string
	format := ".mkv"
	if len(file) > 0 {
		path = file[0]
		if !strings.HasSuffix(path, format) {
			path += format
		}
	} else {
		path = fmt.Sprintf("%s-%s%s", l.Info.Owner.Username, time.Unix(l.Info.CreateTime, 0).Format("2006y01m02dT15h04m05s"), format)
	}
	if _, err := os.Stat(path); err == nil {
		t := strings.TrimSuffix(path, format)
		path = fmt.Sprintf("%s-%d%s", t, time.Now().Unix(), format)
	}

	// Run ffmpeg command
	cmd := exec.Command("ffmpeg", "-i", url, "-c", "copy", path)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	mu := new(sync.Mutex)
	finished := false

	go func(c *exec.Cmd) {
		<-l.done()

		mu.Lock()
		defer mu.Unlock()
		if !finished {
			// Send q key press to quit
			stdin.Write([]byte("q\n"))
		}
	}(cmd)

	// Go routine to wait for process to exit and return result
	l.t.wg.Add(1)
	go func(c *exec.Cmd, stdout, stderr io.ReadCloser) {
		defer l.t.wg.Done()

		stdoutb, _ := io.ReadAll(stdout)
		stderrb, _ := io.ReadAll(stderr)

		if err := cmd.Wait(); err != nil {
			nerr := new(exec.ExitError)
			if errors.As(err, &nerr) {
				l.t.errHandler(fmt.Errorf("Download command failed with: %w\nCommand: %v\nStderr: %v\nStdout: %v\n", err, cmd.Args, string(stderrb), string(stdoutb)))

			}
			l.t.errHandler(fmt.Errorf("Download command failed with: %w\nCommand: %v\n", err, cmd.Args))
		}

		mu.Lock()
		defer mu.Unlock()
		finished = true
		l.t.infoHandler(fmt.Sprintf("Download for %s finished!", l.Info.Owner.Username))
	}(cmd, stdout, stderr)

	l.t.infoHandler(fmt.Sprintf("Started downloading stream by %s to %s\n", l.Info.Owner.Username, path))

	return nil
}

// Only able to get this while logged in
// func (l *Live) GetRankList() (*RankList, error) {
// 	t := l.t
//
// 	params := copyMap(defaultGETParams)
// 	params["room_id"] = l.ID
// 	params["channel"] = "tiktok_web"
// 	params["anchor_id"] = "idk"
//
// 	body, err := t.sendRequest(&reqOptions{
// 		Endpoint: urlRankList,
// 		Query:    params,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var rsp rankListRsp
// 	if err := json.Unmarshal(body, &rsp); err != nil {
// 		return nil, err
// 	}
//
// 	return &rsp.RankList, nil
// }
