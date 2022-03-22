package gotiktoklive

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	pb "github.com/Davincible/gotiktoklive/proto"

	"context"

	"google.golang.org/protobuf/proto"
)

const (
	POLLING_INTERVAL = time.Second
)

type Live struct {
	tiktok *TikTok

	cursor   string
	wss      net.Conn
	wsURL    string
	wsParams map[string]string

	ID       string
	Info     *RoomInfo
	GiftInfo *GiftInfo
	Events   chan interface{}
}

func (t *TikTok) TrackLive(username string) (*Live, error) {
	id, err := t.getRoomID(username)
	if err != nil {
		return nil, err
	}

	live := Live{
		tiktok: t,
		ID:     id,
		Events: make(chan interface{}, 100),
	}

	roomInfo, err := live.getRoomInfo()
	if err != nil {
		return nil, err
	}
	live.Info = roomInfo

	giftInfo, err := live.getGiftInfo()
	if err != nil {
		return nil, err
	}
	live.GiftInfo = giftInfo

	err = live.getRoomData()
	if err != nil {
		return nil, err
	}

	wss, err := live.tryConnectionUpgrade()
	if err != nil {
		return nil, err
	}
	if !wss {
		live.startPolling(context.Background())
	}

	return &live, nil
}

func (t *TikTok) getRoomID(user string) (string, error) {
	body, err := t.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf(urlUserLive, user),
		OmitAPI:  true,
	})
	if err != nil {
		return "", err
	}

	if id := reRoomIDMeta.FindSubmatch(body); len(id) > 1 {
		return string(id[1]), nil
	}

	if id := reRoomIDJson.FindSubmatch(body); len(id) > 1 {
		return string(id[1]), nil
	}

	if strings.Contains(string(body), `"og:url"`) {
		return "", ErrUserOffline
	}
	return "", ErrIPBlocked
}

func (l *Live) getRoomInfo() (*RoomInfo, error) {
	t := l.tiktok

	params := defaultGETParams
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
	t := l.tiktok

	params := defaultGETParams
	params["room_id"] = l.ID

	body, err := t.sendRequest(&reqOptions{
		Endpoint: urlGiftInfo,
		Query:    params,
	})
	if err != nil {
		return nil, err
	}

	var rsp GiftInfoRsp
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return rsp.GiftInfo, nil
}

func (l *Live) getRoomData() error {
	t := l.tiktok

	params := defaultGETParams
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

func (l *Live) startPolling(ctx context.Context) {
	ticker := time.NewTicker(POLLING_INTERVAL)
	defer ticker.Stop()

	if l.tiktok.Debug {
		l.tiktok.debugHandler("Started polling")
	}

	for {
		select {
		case <-ticker.C:
			err := l.getRoomData()
			if err != nil {
				handleRoutineError(err)
			}

			wss, err := l.tryConnectionUpgrade()
			if err != nil {
				handleRoutineError(err)
			}
			if wss {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
