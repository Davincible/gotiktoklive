package gotiktoklive

import (
	"encoding/base64"
	"fmt"
	pb "gotiktoklive/proto"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func parseMsg(msg *pb.Message, warnHandler func(...interface{})) (protoreflect.ProtoMessage, error) {
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
	case "WebcastLiveIntroMessage":
		out = &pb.WebcastLiveIntroMessage{}
	case "WebcastInRoomBannerMessage":
		out = &pb.WebcastInRoomBannerMessage{}
	default:
		data := base64.StdEncoding.EncodeToString(msg.Binary)
		warnHandler(fmt.Errorf("%w: %s,\n%s", ErrMsgNotImplemented, msg.Type, data))
		return nil, nil
	}
	if err := proto.Unmarshal(msg.Binary, out); err != nil {
		return nil, err
	}
	return out, nil
}

func defaultHandler(i ...interface{}) {
	fmt.Println(i...)
}
