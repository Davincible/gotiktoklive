package gotiktoklive

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	pb "gotiktoklive/proto"

	"google.golang.org/protobuf/proto"
)

type Live struct {
	tiktok *TikTok

	cursor string
	wss    net.Conn

	ID       string
	Info     *RoomInfo
	GiftInfo *GiftInfo
	Events   chan interface{}
}

func (t *TikTok) TrackLive(username string) (*Live, error) {
	live := Live{tiktok: t}

	id, err := t.getRoomID(username)
	if err != nil {
		return nil, err
	}

	live.ID = id
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

func (l *Live) getRoomData() (*pb.WebcastResponse, error) {
	t := l.tiktok

	params := defaultGETParams
	params["room_id"] = l.ID

	body, err := t.sendRequest(&reqOptions{
		Endpoint: urlRoomData,
		Query:    params,
	})
	if err != nil {
		return nil, err
	}
	var rsp pb.WebcastResponse
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	l.cursor = rsp.Cursor
	return &rsp, nil
}
