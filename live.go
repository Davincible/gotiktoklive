package gotiktoklive

import (
	"encoding/json"
	"net"
	"time"

	pb "github.com/Davincible/gotiktoklive/proto"
	"golang.org/x/net/context"

	"google.golang.org/protobuf/proto"
)

// TODO: check gift prices of gifts not in wish list

const (
	POLLING_INTERVAL = time.Second
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
}

func (t *TikTok) newLive(roomId string) *Live {
	live := Live{
		t:      t,
		ID:     roomId,
		Events: make(chan interface{}, 100),
	}

	ctx, cancel := context.WithCancel(context.Background())
	live.done = ctx.Done
	live.close = func() {
		cancel()
		close(live.Events)
	}

	return &live
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

// TrackUser will start to track the livestream of a user, if live.
// To listen to events emitted by the livestream, such as comments and viewer
//  count, listen to the Live.Events channel.
func (t *TikTok) TrackUser(username string) (*Live, error) {
	id, err := t.getRoomID(username)
	if err != nil {
		return nil, err
	}

	return t.TrackRoom(id)
}

// TrackRoom will start to track a room by room ID
func (t *TikTok) TrackRoom(roomId string) (*Live, error) {
	live := t.newLive(roomId)

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

	if l.t.Debug {
		l.t.debugHandler("Started polling")
	}

	for {
		select {
		case <-ticker.C:
			err := l.getRoomData()
			if err != nil {
				l.t.errHandler(err)
			}

			wss, err := l.tryConnectionUpgrade()
			if err != nil {
				l.t.errHandler(err)
			}
			if wss {
				return
			}
		case <-l.t.done():
			return
		}
	}
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
