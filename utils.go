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
					Event: toUserType(pt.Event.EventDetails.DisplayType),
					User:  toUser(pt.User),
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
				Event: toUserType(pt.Event.EventDetails.DisplayType),
				User:  toUser(pt.User),
			}
		}()
	case "WebcastGiftMessage":
		pt = &pb.WebcastGiftMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastGiftMessage)

			out = GiftEvent{
				ID:          int(pt.GiftId),
				Name:        pt.GiftDetails.GiftName,
				Describe:    pt.GiftDetails.Describe,
				Cost:        int(pt.GiftDetails.DiamondCount),
				RepeatCount: int(pt.RepeatCount),
				RepeatEnd:   pt.RepeatEnd == 1,
				Type:        int(pt.GiftDetails.GiftType),
				ToUserID:    int64(pt.GiftExtra.ToUserId),
				Timestamp:   int64(pt.GiftExtra.Timestamp),
				User:        toUser(pt.User),
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
			users := []*User{}
			for _, u := range pt.BattleUsers {
				user := u.BattleGroup.User

				users = append(users, &User{
					ID:       int64(user.UserId),
					Username: user.UniqueId,
					Nickname: user.Nickname,
					ProfilePicture: &ProfilePicture{
						Urls: user.ProfilePicture.Urls,
					},
				})
			}
			out = MicBattleEvent{
				Users: users,
			}
		}()
	case "WebcastLinkMicArmies":
		pt = &pb.WebcastLinkMicArmies{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastLinkMicArmies)
			battles := []*Battle{}
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
				battles = append(battles, battle)
			}
			out = BattlesEvent{
				Status:  int(pt.BattleStatus),
				Battles: battles,
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
				ID:    int(pt.Id),
				Title: pt.Description,
				User:  toUser(pt.User),
			}
		}()
	case "WebcastInRoomBannerMessage":
		pt = &pb.WebcastInRoomBannerMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastInRoomBannerMessage)

			var data interface{}
			err = json.Unmarshal([]byte(pt.Json), &data)
			if err != nil {
				err = fmt.Errorf("WebcastInRoomBannerMessage: %w\n%s", err, pt.Json)
			}

			out = RoomBannerEvent{
				Data: data,
			}
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
	case "WebcastBottomMessage":
		pt = &pb.WebcastBottomMessage{}
		defer func() {
			if err != nil {
				return
			}

			pt := pt.(*pb.WebcastBottomMessage)
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

		// Unimplemented Events. Examples can be decoded at : https://protobuf-decoder.netlify.app/
	case "WebcastEnvelopeMessage":
		// Example: Ci4KFldlYmNhc3RFbnZlbG9wZU1lc3NhZ2UQhZab7qCftKViGIGWgM7G8KylYjABEjIKEzcwODI2ODgxODIxMzU1ODk2MzgaBm1hbGl2YVoTNzA4MjY3MDc0NTI4NDc3NDY1NxgC
		return nil, nil
	case "WebcastGiftBroadcastMessage":
		// Example: CkQKG1dlYmNhc3RHaWZ0QnJvYWRjYXN0TWVzc2FnZRCulrKEpuy0pWIYgpaFyNGiraViMAGKAQ5naWZ0X2V4cGVuc2l2ZRCFiM6wwIXJ8WAaiAIKZ2h0dHBzOi8vcDE2LXdlYmNhc3QudGlrdG9rY2RuLmNvbS9pbWcvbWFsaXZhL3dlYmNhc3QtdmEvOTZmMjk3MDYyNGE3OWNkNGQ5NWJlNzI4NDQ5ZDVjODl+dHBsdi1vYmouaW1hZ2UKZ2h0dHBzOi8vcDE5LXdlYmNhc3QudGlrdG9rY2RuLmNvbS9pbWcvbWFsaXZhL3dlYmNhc3QtdmEvOTZmMjk3MDYyNGE3OWNkNGQ5NWJlNzI4NDQ5ZDVjODl+dHBsdi1vYmouaW1hZ2USK3dlYmNhc3QtdmEvOTZmMjk3MDYyNGE3OWNkNGQ5NWJlNzI4NDQ5ZDVjODkqByNCMUNDQTMi/hAK1A4KGFdlYmNhc3RSb29tTm90aWZ5TWVzc2FnZRCulrKEpuy0pWIYgpaFyNGiraViIM+Vp6L/LzABQpoOCiVwbV9tdF9saXZlX2dpZnRfcGxhdGZvcm1fYW5ub3VuY2VtZW50EiJ7MDp1c2VyfSBzZW50IHsxOmdpZnR9IHRvIHsyOnVzZXJ9Gg4KCSNmZmZmZmZmZiCQAyKfBwgLqgGZBwqWBwiFiM6wwIXJ8WAaDVNvbGRpZXJHdXJsODJKjAYKugFodHRwczovL3AxNi1zaWduLnRpa3Rva2Nkbi11cy5jb20vdG9zLXVzZWFzdDUtYXZ0LTAwNjgtdHgvZjEzODBiMTY0MTBmZTFkMzY0MjBkZmRjMDAzYTY0NmJ+dHBsdi10aWt0b2stc2hyaW5rOjcyOjcyLndlYnA/eC1leHBpcmVzPTE2NDkxNTY0MDAmeC1zaWduYXR1cmU9QUY3JTJGeWc3VlNXZFAxT09KR1F6ZFBWM1UwS3MlM0QKrAFodHRwczovL3AxNi1zaWduLnRpa3Rva2Nkbi11cy5jb20vdG9zLXVzZWFzdDUtYXZ0LTAwNjgtdHgvZjEzODBiMTY0MTBmZTFkMzY0MjBkZmRjMDAzYTY0NmJ+YzVfMTAweDEwMC53ZWJwP3gtZXhwaXJlcz0xNjQ5MTU2NDAwJngtc2lnbmF0dXJlPUl4N3NZM2k2ZXolMkJjMVFhYWwyNGxQdHZrYm5jJTNECqwBaHR0cHM6Ly9wMTktc2lnbi50aWt0b2tjZG4tdXMuY29tL3Rvcy11c2Vhc3Q1LWF2dC0wMDY4LXR4L2YxMzgwYjE2NDEwZmUxZDM2NDIwZGZkYzAwM2E2NDZifmM1XzEwMHgxMDAud2VicD94LWV4cGlyZXM9MTY0OTE1NjQwMCZ4LXNpZ25hdHVyZT1FSFNwU01JZFNwMU1GdCUyRmo0cWozcnJDdnpSayUzRAqsAWh0dHBzOi8vcDE2LXNpZ24udGlrdG9rY2RuLXVzLmNvbS90b3MtdXNlYXN0NS1hdnQtMDA2OC10eC9mMTM4MGIxNjQxMGZlMWQzNjQyMGRmZGMwMDNhNjQ2Yn5jNV8xMDB4MTAwLmpwZWc/eC1leHBpcmVzPTE2NDkxNTY0MDAmeC1zaWduYXR1cmU9cWt0c3JsTGQlMkJ1d0J4ZVl1d0hNakN4NTFRTkElM0QSQDEwMHgxMDAvdG9zLXVzZWFzdDUtYXZ0LTAwNjgtdHgvZjEzODBiMTY0MTBmZTFkMzY0MjBkZmRjMDAzYTY0NmKyAQYIgQ8Qwhm6AQCCAgCyAg1zb2xkaWVyZ3VybDgy8gJMTVM0d0xqQUJBQUFBM1dKX0pRWGtpV01jaDBPeW92a3pVRUVfMXZvM1ZDU1ptdVFiRTBzNVVSSG1mLUM2YVdxOERPemZISktLNWtjeiItCAyyASgIli8SIQoObGl2ZV9naWZ0XzYwMzgSD1Rpa1RvayBVbml2ZXJzZRgBIusFCAuqAeUFCuIFCIWI4KSi4efZXhoMU2hhbmUgbGl0dGxlStMECrkBaHR0cHM6Ly9wMTYtc2lnbi1zZy50aWt0b2tjZG4uY29tL3Rvcy1hbGlzZy1hdnQtMDA2OC80MzNkZjRjZmJiMTE2OTliNmU3OGNlN2Q2ZDlhY2U1Nn50cGx2LXRpa3Rvay1zaHJpbms6NzI6NzIud2VicD94LWV4cGlyZXM9MTY0OTE1NjQwMCZ4LXNpZ25hdHVyZT0wckQzSGlWWSUyRlYlMkJmQSUyQm1FYmR6ZHgyUUphZUUlM0QKrAFodHRwczovL3AxNi1zaWduLXNnLnRpa3Rva2Nkbi5jb20vYXdlbWUvMTAweDEwMC90b3MtYWxpc2ctYXZ0LTAwNjgvNDMzZGY0Y2ZiYjExNjk5YjZlNzhjZTdkNmQ5YWNlNTYud2VicD94LWV4cGlyZXM9MTY0OTE1NjQwMCZ4LXNpZ25hdHVyZT01YnVHMk4yQUNTU0VSQnVPUkpUdDIlMkJLYUolMkZ3JTNECqgBaHR0cHM6Ly9wMTYtc2lnbi1zZy50aWt0b2tjZG4uY29tL2F3ZW1lLzEwMHgxMDAvdG9zLWFsaXNnLWF2dC0wMDY4LzQzM2RmNGNmYmIxMTY5OWI2ZTc4Y2U3ZDZkOWFjZTU2LmpwZWc/eC1leHBpcmVzPTE2NDkxNTY0MDAmeC1zaWduYXR1cmU9NU1QTjVodWFnWnVwODdqQVI1cm15dlF4empZJTNEEjsxMDB4MTAwL3Rvcy1hbGlzZy1hdnQtMDA2OC80MzNkZjRjZmJiMTE2OTliNmU3OGNlN2Q2ZDlhY2U1NrIBBwjuAhCiswW6AQCCAgCyAhJub3RvcmlvdXNfcC5pLmdfX1/yAkxNUzR3TGpBQkFBQUFDWlRTcE91bHlvUnVuRjBIOFM4em1yTURuY1AzeWw3OEVkeVo4R254S0tOVmpkVDJKcXNlY3ZZZFJ5NmRXTW5KEipzc2xvY2FsOi8vd2ViY2FzdF9naWZ0X2RpYWxvZz9naWZ0X2lkPTYwMzgYAzLmAQgFEuEBCOMCEBgaW2h0dHBzOi8vcDE2LXdlYmNhc3QudGlrdG9rY2RuLmNvbS9pbWcvYWxpc2cvd2ViY2FzdC1zZy9icm9hZGNhc3RfZ2lmdF9iZy5wbmd+dHBsdi1vYmouaW1hZ2UaW2h0dHBzOi8vcDE5LXdlYmNhc3QudGlrdG9rY2RuLmNvbS9pbWcvYWxpc2cvd2ViY2FzdC1zZy9icm9hZGNhc3RfZ2lmdF9iZy5wbmd+dHBsdi1vYmouaW1hZ2UiIHdlYmNhc3Qtc2cvYnJvYWRjYXN0X2dpZnRfYmcucG5nSg5naWZ0X2Jyb2FkY2FzdA==
		return nil, nil
	case "WebcastLinkmicBattleNoticeMessage":
		// Example: CkAKIVdlYmNhc3RMaW5rbWljQmF0dGxlTm90aWNlTWVzc2FnZRCClq381661pWIYgZaAzsbwrKViILyUyKL/LygBEAQ6fAoiCiBMZWFybiB0byBwbGF5IGFuZCB3aW4gdGhlIG1hdGNoIRIGCgRWaWV3GkxodHRwczovL3dlYmNhc3QudGlrdG9rdi5jb20vZmFsY29uL3dlYmNhc3RfbXQvcGFnZS9saXZlX21hdGNoL2ZhcS9pbmRleC5odG1sIAE=
		return nil, nil
	case "WebcastHourlyRankMessage":
		// Example: CjUKGFdlYmNhc3RIb3VybHlSYW5rTWVzc2FnZRCFlo/+3bm1pWIYgZaAzsbwrKViIL6S0aL/LxI1CAEiLwoNcG1fbXRfTGl2ZV9XUhIOV2Vla2x5IHJhbmtpbmcaDgoJI2ZmZmZmZmZmIJADMAEYAQ==
		return nil, nil
	case "LinkMicMethod":
		// Example: CiwKDUxpbmtNaWNNZXRob2QQgZaP+unbtaViGIGWgM7G8KylYiDLjN6i/y8oARAEQIKWyO7J9qylYkgEUATYAgI=
		return nil, nil
	case "WebcastLinkMessage":
		// Example: Ci8KEldlYmNhc3RMaW5rTWVzc2FnZRCBlpqQ/dy1pWIYgZaAzsbwrKViIOCL3qL/LxACGIKWyO7J9qylYiAC
		return nil, nil
	case "WebcastLinkMicBattlePunishFinish":
		// Example: Cj8KIFdlYmNhc3RMaW5rTWljQmF0dGxlUHVuaXNoRmluaXNoEIKWiaD+2rWlYhiBloDOxvCspWIgxYreov8vKAEQgpbI7sn2rKViGIGIuJ6L2KnnYSABKIKWoYrLrLWlYg==
		return nil, nil
	case "WebcastUnauthorizedMemberMessage":
		// Example: Cj8KIFdlYmNhc3RVbmF1dGhvcml6ZWRNZW1iZXJNZXNzYWdlEIGWipDYoOSmYhiBlon8wPXdpmIgi//59/8vMAEQARojChF3ZWJfbm9ubG9naW5faW1fMRoOCgkjZmZmZmZmZmYgkAMiBjc0MDQzNCo4ChVsaXZlX3Jvb21fZW50ZXJfdG9hc3QSD3swOnVzZXJ9IGpvaW5lZBoOCgkjZmZmZmZmZmYgkAM=
		// TODO: do something here
		return nil, nil
	case "WebcastRankUpdateMessage":
		// Example: CjUKGFdlYmNhc3RSYW5rVXBkYXRlTWVzc2FnZRCBloyOuZSKqmIYgZaGnIDeiapiIPre/MWBMBJECAEaLwoNcG1fbXRfTGl2ZV9XUhIOV2Vla2x5IHJhbmtpbmcaDgoJI2ZmZmZmZmZmIJADIgsiCSM4MEZGMzY3RjDA4yM=
		return nil, nil
	case "WebcastLinkMicMethod":
		// Example: CjMKFFdlYmNhc3RMaW5rTWljTWV0aG9kEK6WofTLnoCrYhiBlqyCgfz7qmIgiLPQ/4EwMAEQCCiBiIDurevuuF8wuSs4uSs=
		return nil, nil
	case "WebcastRankTextMessage":
		// Used to broadcast the rank of a user, i.e. a top donator
		// Example: Cv8GChZXZWJjYXN0UmFua1RleHRNZXNzYWdlEK6Wjsyclq+rYhialonA1vCoq2IgjILSloIwQskGChdwbV9tdF90b3B2aWV3ZXJfY29tbWVudBItezA6dXNlcn0ganVzdCBiZWNhbWUgYSB0b3AgezE6c3RyaW5nfSB2aWV3ZXIhGg4KCSNmZmZmZmZmZiCQAyLnBQgLqgHhBQreBQiFiMra5PGp2l4aFkpVTElFIOKdpO+4j/CfjJ7inaTvuI9KzQQKqgFodHRwczovL3AxOS1zaWduLnRpa3Rva2Nkbi11cy5jb20vdG9zLXVzZWFzdDUtYXZ0LTAwNjgtdHgvMzQzNmNkZTE2OWVhOTE5Nzg2NTUwYzNkOWNlOWU0OTB+YzVfMTAweDEwMC53ZWJwP3gtZXhwaXJlcz0xNjQ5OTM3NjAwJngtc2lnbmF0dXJlPVE4SUJzbUkzM21YRDZjeUg3dGc2Mlh1TlppcyUzRAquAWh0dHBzOi8vcDE2LXNpZ24udGlrdG9rY2RuLXVzLmNvbS90b3MtdXNlYXN0NS1hdnQtMDA2OC10eC8zNDM2Y2RlMTY5ZWE5MTk3ODY1NTBjM2Q5Y2U5ZTQ5MH5jNV8xMDB4MTAwLndlYnA/eC1leHBpcmVzPTE2NDk5Mzc2MDAmeC1zaWduYXR1cmU9JTJGNGRTcWowJTJGeDZCWDhIcWoxRXQ3RWh3d0x0USUzRAqqAWh0dHBzOi8vcDE5LXNpZ24udGlrdG9rY2RuLXVzLmNvbS90b3MtdXNlYXN0NS1hdnQtMDA2OC10eC8zNDM2Y2RlMTY5ZWE5MTk3ODY1NTBjM2Q5Y2U5ZTQ5MH5jNV8xMDB4MTAwLmpwZWc/eC1leHBpcmVzPTE2NDk5Mzc2MDAmeC1zaWduYXR1cmU9NjZZdVU4d3pqbkplN252ODJCM3BseXpHZ3RnJTNEEkAxMDB4MTAwL3Rvcy11c2Vhc3Q1LWF2dC0wMDY4LXR4LzM0MzZjZGUxNjllYTkxOTc4NjU1MGMzZDljZTllNDkwsgEGCJgREOoSggIAsgIOanVsaWViZXlvdTExMTHyAkxNUzR3TGpBQkFBQUFVR0FFQVRpemlCd2swQ3d0c1dzcy1TX1pVM1dBTUNtbVNkYm9GazdRbmFiZU5FYU5Cb3R4aEIzQl9GeVo4N0RoIgUIAVoBMxABGAQgAw==
		return nil, nil
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
	if u == nil {
		return &User{}
	}

	user := User{
		ID:       int64(u.UserId),
		Username: u.UniqueId,
		Nickname: u.Nickname,
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

func toUserType(displayType string) userEventType {
	switch displayType {
	case "pm_main_follow_message_viewer_2":
		return USER_FOLLOW
	case "pm_mt_guidance_share":
		return USER_SHARE
	case "live_room_enter_toast":
		return USER_JOIN
	}
	return userEventType(fmt.Sprintf("User type not implemented, please report: %s", displayType))
}
