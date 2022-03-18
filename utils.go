package gotiktoklive

import (
	"fmt"
	pb "gotiktoklive/proto"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func parseMsg(msg *pb.Message) (protoreflect.ProtoMessage, error) {
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
