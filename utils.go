package gotiktoklive

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	pb "github.com/Davincible/gotiktoklive/proto"

	"google.golang.org/protobuf/proto"
)

func parseMsg(msg *pb.Message, warnHandler func(...interface{})) (out interface{}, err error) {
	var pt proto.Message
	switch msg.Type {
	case "WebcastChatMessage":
		pt = &pb.WebcastChatMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastChatMessage)
			out = ChatEvent{
				Comment: pt.Comment,
				User:    toUser(pt.User),
			}
		}()
	case "WebcastMemberMessage":
		// UserEvent
		pt = &pb.WebcastMemberMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastMemberMessage)
			if pt.Event != nil && pt.Event.EventDetails != nil {
				out = UserEvent{
					DisplayType: pt.Event.EventDetails.DisplayType,
					Label:       pt.Event.EventDetails.Label,
					User:        toUser(pt.User),
				}
			}
		}()
	case "WebcastRoomUserSeqMessage":
		pt = &pb.WebcastRoomUserSeqMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastRoomUserSeqMessage)
			out = ViewersEvent{
				Viewers: int(pt.ViewerCount),
			}
		}()
	case "WebcastSocialMessage":
		pt = &pb.WebcastSocialMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastSocialMessage)
			out = UserEvent{
				DisplayType: pt.Event.EventDetails.DisplayType,
				Label:       pt.Event.EventDetails.Label,
				User:        toUser(pt.User),
			}
		}()
	case "WebcastGiftMessage":
		pt = &pb.WebcastGiftMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastGiftMessage)

			var gift Gift
			err = json.Unmarshal([]byte(pt.GiftJson), &gift)

			out = GiftEvent{
				Gift: gift,
				User: toUser(pt.User),
			}
		}()
	case "WebcastLikeMessage":
		pt = &pb.WebcastLikeMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastLikeMessage)
			out = LikeEvent{
				Likes:       int(pt.LikeCount),
				TotalLikes:  int(pt.TotalLikeCount),
				User:        toUser(pt.User),
				DisplayType: pt.Event.EventDetails.DisplayType,
				Label:       pt.Event.EventDetails.Label,
			}
		}()
	case "WebcastQuestionNewMessage":
		pt = &pb.WebcastQuestionNewMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastQuestionNewMessage)
			out = QuestionEvent{
				Quesion: pt.QuestionDetails.QuestionText,
				User:    toUser(pt.QuestionDetails.User),
			}
		}()
	case "WebcastWebsocketMessage":
		pt = &pb.WebcastWebsocketMessage{}
		// TODO: implement
	case "WebcastControlMessage":
		pt = &pb.WebcastControlMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastControlMessage)
			out = ControlEvent{
				Action: int(pt.Action),
			}
		}()
	case "WebcastLinkMicBattle":
		pt = &pb.WebcastLinkMicBattle{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastLinkMicBattle)
			out = MicBattleEvent{
				Users: []*User{},
			}
			for _, u := range pt.BattleUsers {
				user := u.BattleGroup.User
				p := out.(*MicBattleEvent)
				p.Users = append(p.Users, &User{
					ID:       int64(user.UserId),
					Nickname: user.Nickname,
					FullName: user.UniqueId,
					ProfilePicture: &ProfilePicture{
						Urls: user.ProfilePicture.Urls,
					},
				})
			}
		}()
	case "WebcastLinkMicArmies":
		pt = &pb.WebcastLinkMicArmies{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastLinkMicArmies)
			out = BattlesEvent{
				Status:  int(pt.BattleStatus),
				Battles: []*Battle{},
			}
			for _, b := range pt.BattleItems {
				battle := &Battle{
					Host:   int64(b.HostUserId),
					Groups: []*BattleGroup{},
				}
				for _, g := range b.BattleGroups {
					group := BattleGroup{
						Points: int(g.Points),
						Users:  []*User{},
					}
					for _, u := range g.Users {
						group.Users = append(group.Users, toUser(u))
					}
					battle.Groups = append(battle.Groups, &group)
				}
				p := out.(*BattlesEvent)
				p.Battles = append(p.Battles, battle)
			}
		}()
	case "WebcastLiveIntroMessage":
		pt = &pb.WebcastLiveIntroMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastLiveIntroMessage)
			out = IntroEvent{
				ID:          int(pt.Id),
				Description: pt.Description,
				User:        toUser(pt.User),
			}
		}()
	case "WebcastInRoomBannerMessage":
		pt = &pb.WebcastInRoomBannerMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastInRoomBannerMessage)
			out = RoomBannerEvent{}
			err = json.Unmarshal([]byte(pt.Json), &out.(*RoomBannerEvent).Data)
		}()
	case "RoomMessage":
		pt = &pb.RoomMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.RoomMessage)
			out = RoomEvent{
				Type:    pt.Type.Type,
				Message: pt.Text,
			}
		}()
	case "WebcastWishlistUpdateMessage":
		pt = &pb.WebcastWishlistUpdateMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastWishlistUpdateMessage)
			out = pt
		}()
	default:
		data := base64.StdEncoding.EncodeToString(msg.Binary)
		warnHandler(fmt.Errorf("%w: %s,\n%s", ErrMsgNotImplemented, msg.Type, data))
		return nil, nil
	}
	if err = proto.Unmarshal(msg.Binary, pt); err != nil {
		base := base64.RawStdEncoding.EncodeToString(msg.Binary)
		err = fmt.Errorf("Failed to unmarshal proto %T: %w\n%s", pt, err, base)
	}
	return
}

func defaultLogHandler(i ...interface{}) {
	fmt.Println(i...)
}

func routineErrHandler(err error) {
	panic(err)
}

func toUser(u *pb.User) *User {
	user := User{
		ID:       int64(u.UserId),
		Nickname: u.Nickname,
		FullName: u.UniqueId,
	}

	if u.ProfilePicture != nil && u.ProfilePicture.Urls != nil {
		user.ProfilePicture = &ProfilePicture{
			Urls: u.ProfilePicture.Urls,
		}
	}

	if u.ExtraAttributes != nil {
		user.ExtraAttributes = &ExtraAttributes{
			FollowRole: int(u.ExtraAttributes.FollowRole),
		}
	}

	if u.Badge != nil && u.Badge.Badges != nil {
		var badges []*UserBadge
		for _, badge := range u.Badge.Badges {
			badges = append(badges, &UserBadge{
				Type: badge.Type,
				Name: badge.Name,
			})
		}
		user.Badge = &BadgeAttributes{
			Badges: badges,
		}
	}
	return &user
}

func copyMap(m map[string]string) map[string]string {
	out := make(map[string]string)
	for key, value := range m {
		out[key] = value
	}
	return out
}
