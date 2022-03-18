package gotiktoklive

import (
	"encoding/json"
	"fmt"
	pb "gotiktoklive/proto"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type TikTok struct {
	c *http.Client

	Debug bool

	infoHandler  func(...interface{})
	warnHandler  func(...interface{})
	debugHandler func(...interface{})
}

func NewTikTok() *TikTok {
	return &TikTok{
		c:            &http.Client{},
		infoHandler:  defaultHandler,
		warnHandler:  defaultHandler,
		debugHandler: defaultHandler,
	}
}

func GetRoomInfo(user string) error {
	return nil
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

func (t *TikTok) getRoomInfo(id string) (*RoomInfo, error) {
	params := defaultGETParams
	params["room_id"] = id

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

func (t *TikTok) getGiftInfo(id string) (*GiftInfo, error) {
	params := defaultGETParams
	params["room_id"] = id

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

func (t *TikTok) getRoomData(id string) (*pb.WebcastResponse, error) {
	params := defaultGETParams
	params["room_id"] = id

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
	return &rsp, nil
}

func (t *TikTok) parseMsg(msg *pb.Message) (protoreflect.ProtoMessage, error) {
	var out protoreflect.ProtoMessage
	switch msg.Type {
	case "WebcastChatMessage":
		out = &pb.WebcastChatMessage{}
	case "WebcastMemberMessage":
		out = &pb.WebcastMemberMessage{}
	case "WebcastRoomUserSeqMessage":
		out = &pb.WebcastRoomUserSeqMessage{}
	case "WebcastSocialMessage":
		out = &pb.WebcastSocialMessage{}
	case "WebcastGiftMessage":
		out = &pb.WebcastGiftMessage{}
	case "WebcastLikeMessage":
		out = &pb.WebcastLikeMessage{}
	case "WebcastQuestionNewMessage":
		out = &pb.WebcastQuestionNewMessage{}
	case "WebcastWebsocketMessage":
		out = &pb.WebcastWebsocketMessage{}
	case "WebcastControlMessage":
		out = &pb.WebcastControlMessage{}
	case "WebcastLinkMicBattle":
		out = &pb.WebcastLinkMicBattle{}
	case "WebcastLinkMicArmies":
		out = &pb.WebcastLinkMicArmies{}
	default:
		return nil, ErrMsgNotImplemented
	}
	if err := proto.Unmarshal(msg.Binary, out); err != nil {
		return nil, err
	}
	return out, nil
}

func defaultHandler(i ...interface{}) {
	fmt.Println(i...)
}
