package gotiktoklive

type roomInfoRsp struct {
	RoomInfo *RoomInfo `json:"data"`
	Extra    struct {
		Now int64 `json:"now"`
	} `json:"extra"`
	StatusCode float64 `json:"status_code"`
}

type RoomInfo struct {
	AnchorABMap struct {
	} `json:"AnchorABMap"`
	AdminUserIds             []interface{} `json:"admin_user_ids"`
	AnchorScheduledTimeText  string        `json:"anchor_scheduled_time_text"`
	AnchorShareText          string        `json:"anchor_share_text"`
	AnchorTabType            float64       `json:"anchor_tab_type"`
	AnsweringQuestionContent string        `json:"answering_question_content"`
	AppID                    float64       `json:"app_id"`
	AutoCover                float64       `json:"auto_cover"`
	BookEndTime              float64       `json:"book_end_time"`
	BookTime                 float64       `json:"book_time"`
	BusinessLive             float64       `json:"business_live"`
	ChallengeInfo            string        `json:"challenge_info"`
	ClientVersion            float64       `json:"client_version"`
	CommentNameMode          float64       `json:"comment_name_mode"`
	CommerceInfo             struct {
		CommercePermission       float64 `json:"commerce_permission"`
		OecLiveEnterRoomInitData string  `json:"oec_live_enter_room_init_data"`
	} `json:"commerce_info"`
	CommonLabelList string `json:"common_label_list"`
	ContentTag      string `json:"content_tag"`
	Cover           struct {
		AvgColor   string   `json:"avg_color"`
		Height     float64  `json:"height"`
		ImageType  float64  `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      float64  `json:"width"`
	} `json:"cover"`
	CreateTime           float64       `json:"create_time"`
	DecoList             []interface{} `json:"deco_list"`
	DisablePreloadStream bool          `json:"disable_preload_stream"`
	FansclubMsgStyle     float64       `json:"fansclub_msg_style"`
	FeedRoomLabel        struct {
		AvgColor   string   `json:"avg_color"`
		Height     float64  `json:"height"`
		ImageType  float64  `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      float64  `json:"width"`
	} `json:"feed_room_label"`
	FeedRoomLabels      []interface{} `json:"feed_room_labels"`
	FilterMsgRules      []interface{} `json:"filter_msg_rules"`
	FinishReason        float64       `json:"finish_reason"`
	FinishTime          float64       `json:"finish_time"`
	FinishURL           string        `json:"finish_url"`
	FinishURLV2         string        `json:"finish_url_v2"`
	FollowMsgStyle      float64       `json:"follow_msg_style"`
	ForumExtraData      string        `json:"forum_extra_data"`
	GameTag             []interface{} `json:"game_tag"`
	GiftMsgStyle        float64       `json:"gift_msg_style"`
	GiftPollVoteEnabled bool          `json:"gift_poll_vote_enabled"`
	GroupSource         float64       `json:"group_source"`
	HasCommerceGoods    bool          `json:"has_commerce_goods"`
	Hashtag             struct {
		ID    float64 `json:"id"`
		Image struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"image"`
		Namespace float64 `json:"namespace"`
		Title     string  `json:"title"`
	} `json:"hashtag"`
	HaveWishlist               bool    `json:"have_wishlist"`
	HotSentenceInfo            string  `json:"hot_sentence_info"`
	ID                         int64   `json:"id"`
	IDStr                      string  `json:"id_str"`
	InteractionQuestionVersion float64 `json:"interaction_question_version"`
	Introduction               string  `json:"introduction"`
	IsGatedRoom                bool    `json:"is_gated_room"`
	IsReplay                   bool    `json:"is_replay"`
	IsShowUserCardSwitch       bool    `json:"is_show_user_card_switch"`
	LastPingTime               float64 `json:"last_ping_time"`
	Layout                     float64 `json:"layout"`
	LikeCount                  float64 `json:"like_count"`
	LinkMic                    struct {
		AudienceIDList []interface{} `json:"audience_id_list"`
		BattleScores   []interface{} `json:"battle_scores"`
		BattleSettings struct {
			BattleID    float64 `json:"battle_id"`
			ChannelID   float64 `json:"channel_id"`
			Duration    float64 `json:"duration"`
			Finished    float64 `json:"finished"`
			MatchType   float64 `json:"match_type"`
			StartTime   float64 `json:"start_time"`
			StartTimeMs float64 `json:"start_time_ms"`
			Theme       string  `json:"theme"`
		} `json:"battle_settings"`
		ChannelID      float64       `json:"channel_id"`
		FollowedCount  float64       `json:"followed_count"`
		LinkedUserList []interface{} `json:"linked_user_list"`
		MultiLiveEnum  float64       `json:"multi_live_enum"`
		RivalAnchorID  float64       `json:"rival_anchor_id"`
		ShowUserList   []interface{} `json:"show_user_list"`
	} `json:"link_mic"`
	LinkerMap struct {
	} `json:"linker_map"`
	LinkmicLayout      float64       `json:"linkmic_layout"`
	LiveDistribution   []interface{} `json:"live_distribution"`
	LiveID             float64       `json:"live_id"`
	LiveReason         string        `json:"live_reason"`
	LiveRoomMode       float64       `json:"live_room_mode"`
	LiveTypeAudio      bool          `json:"live_type_audio"`
	LiveTypeLinkmic    bool          `json:"live_type_linkmic"`
	LiveTypeNormal     bool          `json:"live_type_normal"`
	LiveTypeSandbox    bool          `json:"live_type_sandbox"`
	LiveTypeScreenshot bool          `json:"live_type_screenshot"`
	LiveTypeSocialLive bool          `json:"live_type_social_live"`
	LiveTypeThirdParty bool          `json:"live_type_third_party"`
	LivingRoomAttrs    struct {
		AdminFlag   float64 `json:"admin_flag"`
		Rank        float64 `json:"rank"`
		RoomID      int64   `json:"room_id"`
		RoomIDStr   string  `json:"room_id_str"`
		SilenceFlag float64 `json:"silence_flag"`
	} `json:"living_room_attrs"`
	LotteryFinishTime float64 `json:"lottery_finish_time"`
	MosaicStatus      float64 `json:"mosaic_status"`
	OsType            float64 `json:"os_type"`
	Owner             struct {
		AllowFindByContacts                 bool `json:"allow_find_by_contacts"`
		AllowOthersDownloadVideo            bool `json:"allow_others_download_video"`
		AllowOthersDownloadWhenSharingVideo bool `json:"allow_others_download_when_sharing_video"`
		AllowShareShowProfile               bool `json:"allow_share_show_profile"`
		AllowShowInGossip                   bool `json:"allow_show_in_gossip"`
		AllowShowMyAction                   bool `json:"allow_show_my_action"`
		AllowStrangeComment                 bool `json:"allow_strange_comment"`
		AllowUnfollowerComment              bool `json:"allow_unfollower_comment"`
		AllowUseLinkmic                     bool `json:"allow_use_linkmic"`
		AvatarLarge                         struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"avatar_large"`
		AvatarMedium struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"avatar_medium"`
		AvatarThumb struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"avatar_thumb"`
		BadgeImageList           []interface{} `json:"badge_image_list"`
		BadgeList                []interface{} `json:"badge_list"`
		BgImgURL                 string        `json:"bg_img_url"`
		BioDescription           string        `json:"bio_description"`
		BlockStatus              float64       `json:"block_status"`
		BorderList               []interface{} `json:"border_list"`
		CommentRestrict          float64       `json:"comment_restrict"`
		CommerceWebcastConfigIds []interface{} `json:"commerce_webcast_config_ids"`
		Constellation            string        `json:"constellation"`
		CreateTime               float64       `json:"create_time"`
		DisableIchat             float64       `json:"disable_ichat"`
		DisplayID                string        `json:"display_id"`
		EnableIchatImg           float64       `json:"enable_ichat_img"`
		Exp                      float64       `json:"exp"`
		FanTicketCount           float64       `json:"fan_ticket_count"`
		FoldStrangerChat         bool          `json:"fold_stranger_chat"`
		FollowInfo               struct {
			FollowStatus   float64 `json:"follow_status"`
			FollowerCount  float64 `json:"follower_count"`
			FollowingCount float64 `json:"following_count"`
			PushStatus     float64 `json:"push_status"`
		} `json:"follow_info"`
		FollowStatus        float64       `json:"follow_status"`
		IchatRestrictType   float64       `json:"ichat_restrict_type"`
		ID                  float64       `json:"id"`
		IDStr               string        `json:"id_str"`
		IsFollower          bool          `json:"is_follower"`
		IsFollowing         bool          `json:"is_following"`
		LinkMicStats        float64       `json:"link_mic_stats"`
		MediaBadgeImageList []interface{} `json:"media_badge_image_list"`
		ModifyTime          float64       `json:"modify_time"`
		NeedProfileGuide    bool          `json:"need_profile_guide"`
		NewRealTimeIcons    []interface{} `json:"new_real_time_icons"`
		Nickname            string        `json:"nickname"`
		OwnRoom             struct {
			RoomIds    []int64  `json:"room_ids"`
			RoomIdsStr []string `json:"room_ids_str"`
		} `json:"own_room"`
		PayGrade struct {
			GradeBanner        string        `json:"grade_banner"`
			GradeDescribe      string        `json:"grade_describe"`
			GradeIconList      []interface{} `json:"grade_icon_list"`
			Level              float64       `json:"level"`
			Name               string        `json:"name"`
			NextName           string        `json:"next_name"`
			NextPrivileges     string        `json:"next_privileges"`
			Score              float64       `json:"score"`
			ScreenChatType     float64       `json:"screen_chat_type"`
			UpgradeNeedConsume float64       `json:"upgrade_need_consume"`
		} `json:"pay_grade"`
		PayScore           float64       `json:"pay_score"`
		PayScores          float64       `json:"pay_scores"`
		PushCommentStatus  bool          `json:"push_comment_status"`
		PushDigg           bool          `json:"push_digg"`
		PushFollow         bool          `json:"push_follow"`
		PushFriendAction   bool          `json:"push_friend_action"`
		PushIchat          bool          `json:"push_ichat"`
		PushStatus         bool          `json:"push_status"`
		PushVideoPost      bool          `json:"push_video_post"`
		PushVideoRecommend bool          `json:"push_video_recommend"`
		RealTimeIcons      []interface{} `json:"real_time_icons"`
		SecUID             string        `json:"sec_uid"`
		Secret             float64       `json:"secret"`
		ShareQrcodeURI     string        `json:"share_qrcode_uri"`
		SpecialID          string        `json:"special_id"`
		Status             float64       `json:"status"`
		TicketCount        float64       `json:"ticket_count"`
		TopFans            []*TopFan     `json:"top_fans"`
		TopVipNo           float64       `json:"top_vip_no"`
		UserAttr           struct {
			IsAdmin      bool    `json:"is_admin"`
			IsMuted      bool    `json:"is_muted"`
			IsSuperAdmin bool    `json:"is_super_admin"`
			MuteDuration float64 `json:"mute_duration"`
		} `json:"user_attr"`
		UserRole                    float64 `json:"user_role"`
		Verified                    bool    `json:"verified"`
		VerifiedContent             string  `json:"verified_content"`
		VerifiedReason              string  `json:"verified_reason"`
		WithCarManagementPermission bool    `json:"with_car_management_permission"`
		WithCommercePermission      bool    `json:"with_commerce_permission"`
		WithFusionShopEntry         bool    `json:"with_fusion_shop_entry"`
	} `json:"owner"`
	OwnerDeviceID        float64 `json:"owner_device_id"`
	OwnerDeviceIDStr     string  `json:"owner_device_id_str"`
	OwnerUserID          float64 `json:"owner_user_id"`
	OwnerUserIDStr       string  `json:"owner_user_id_str"`
	PreEnterTime         float64 `json:"pre_enter_time"`
	PreviewFlowTag       float64 `json:"preview_flow_tag"`
	RanklistAudienceType float64 `json:"ranklist_audience_type"`
	RelationTag          string  `json:"relation_tag"`
	Replay               bool    `json:"replay"`
	RoomAuditStatus      float64 `json:"room_audit_status"`
	RoomAuth             struct {
		Banner              float64 `json:"Banner"`
		BroadcastMessage    float64 `json:"BroadcastMessage"`
		Chat                bool    `json:"Chat"`
		ChatL2              bool    `json:"ChatL2"`
		ChatSubOnly         bool    `json:"ChatSubOnly"`
		Danmaku             bool    `json:"Danmaku"`
		Digg                bool    `json:"Digg"`
		DonationSticker     float64 `json:"DonationSticker"`
		Gift                bool    `json:"Gift"`
		GiftAnchorMt        float64 `json:"GiftAnchorMt"`
		GiftPoll            float64 `json:"GiftPoll"`
		GoldenEnvelope      float64 `json:"GoldenEnvelope"`
		InteractionQuestion bool    `json:"InteractionQuestion"`
		Landscape           float64 `json:"Landscape"`
		LandscapeChat       float64 `json:"LandscapeChat"`
		LuckMoney           bool    `json:"LuckMoney"`
		Poll                float64 `json:"Poll"`
		Promote             bool    `json:"Promote"`
		Props               bool    `json:"Props"`
		PublicScreen        float64 `json:"PublicScreen"`
		QuickChat           float64 `json:"QuickChat"`
		Rank                float64 `json:"Rank"`
		RoomContributor     bool    `json:"RoomContributor"`
		Share               bool    `json:"Share"`
		ShareEffect         float64 `json:"ShareEffect"`
		UserCard            bool    `json:"UserCard"`
		UserCount           float64 `json:"UserCount"`
		Viewers             bool    `json:"Viewers"`
		TransactionHistory  float64 `json:"transaction_history"`
	} `json:"room_auth"`
	RoomCreateAbParam string        `json:"room_create_ab_param"`
	RoomLayout        float64       `json:"room_layout"`
	RoomStickerList   []interface{} `json:"room_sticker_list"`
	RoomTabs          []interface{} `json:"room_tabs"`
	RoomTag           float64       `json:"room_tag"`
	ScrollConfig      string        `json:"scroll_config"`
	SearchID          float64       `json:"search_id"`
	ShareMsgStyle     float64       `json:"share_msg_style"`
	ShareURL          string        `json:"share_url"`
	ShortTitle        string        `json:"short_title"`
	ShortTouchItems   []interface{} `json:"short_touch_items"`
	SocialInteraction struct {
		MultiLive struct {
			UserSettings struct {
				MultiLiveApplyPermission float64 `json:"multi_live_apply_permission"`
			} `json:"user_settings"`
		} `json:"multi_live"`
	} `json:"social_interaction"`
	StartTime float64 `json:"start_time"`
	Stats     struct {
		DiggCount            float64 `json:"digg_count"`
		EnterCount           float64 `json:"enter_count"`
		FanTicket            float64 `json:"fan_ticket"`
		FollowCount          float64 `json:"follow_count"`
		GiftUvCount          float64 `json:"gift_uv_count"`
		ID                   int64   `json:"id"`
		IDStr                string  `json:"id_str"`
		LikeCount            float64 `json:"like_count"`
		ReplayFanTicket      float64 `json:"replay_fan_ticket"`
		ReplayViewers        float64 `json:"replay_viewers"`
		ShareCount           float64 `json:"share_count"`
		TotalUser            float64 `json:"total_user"`
		TotalUserDesp        string  `json:"total_user_desp"`
		UserCountComposition struct {
			MyFollow    float64 `json:"my_follow"`
			Other       float64 `json:"other"`
			VideoDetail float64 `json:"video_detail"`
		} `json:"user_count_composition"`
		Watermelon float64 `json:"watermelon"`
	} `json:"stats"`
	Status      float64       `json:"status"`
	StickerList []interface{} `json:"sticker_list"`
	StreamID    int64         `json:"stream_id"`
	StreamIDStr string        `json:"stream_id_str"`
	StreamURL   struct {
		CandidateResolution []string      `json:"candidate_resolution"`
		CompletePushUrls    []interface{} `json:"complete_push_urls"`
		DefaultResolution   string        `json:"default_resolution"`
		Extra               struct {
			AnchorInteractProfile   float64 `json:"anchor_interact_profile"`
			AudienceInteractProfile float64 `json:"audience_interact_profile"`
			BframeEnable            bool    `json:"bframe_enable"`
			BitrateAdaptStrategy    float64 `json:"bitrate_adapt_strategy"`
			Bytevc1Enable           bool    `json:"bytevc1_enable"`
			DefaultBitrate          float64 `json:"default_bitrate"`
			Fps                     float64 `json:"fps"`
			GopSec                  float64 `json:"gop_sec"`
			HardwareEncode          bool    `json:"hardware_encode"`
			Height                  float64 `json:"height"`
			MaxBitrate              float64 `json:"max_bitrate"`
			MinBitrate              float64 `json:"min_bitrate"`
			Roi                     bool    `json:"roi"`
			SwRoi                   bool    `json:"sw_roi"`
			VideoProfile            float64 `json:"video_profile"`
			Width                   float64 `json:"width"`
		} `json:"extra"`
		FlvPullURL struct {
			FullHd1 string `json:"FULL_HD1"`
			Hd1     string `json:"HD1"`
			Sd1     string `json:"SD1"`
			Sd2     string `json:"SD2"`
		} `json:"flv_pull_url"`
		FlvPullURLParams struct {
			FullHd1 string `json:"FULL_HD1"`
			Hd1     string `json:"HD1"`
			Sd1     string `json:"SD1"`
			Sd2     string `json:"SD2"`
		} `json:"flv_pull_url_params"`
		HlsPullURL    string `json:"hls_pull_url"`
		HlsPullURLMap struct {
		} `json:"hls_pull_url_map"`
		HlsPullURLParams string `json:"hls_pull_url_params"`
		ID               int64  `json:"id"`
		IDStr            string `json:"id_str"`
		LiveCoreSdkData  struct {
			PullData struct {
				Options struct {
					DefaultQuality struct {
						Level      float64 `json:"level"`
						Name       string  `json:"name"`
						Resolution string  `json:"resolution"`
						SdkKey     string  `json:"sdk_key"`
						VCodec     string  `json:"v_codec"`
					} `json:"default_quality"`
					Qualities []struct {
						Level      float64 `json:"level"`
						Name       string  `json:"name"`
						Resolution string  `json:"resolution"`
						SdkKey     string  `json:"sdk_key"`
						VCodec     string  `json:"v_codec"`
					} `json:"qualities"`
				} `json:"options"`
				StreamData string `json:"stream_data"`
			} `json:"pull_data"`
		} `json:"live_core_sdk_data"`
		Provider       float64       `json:"provider"`
		PushUrls       []interface{} `json:"push_urls"`
		ResolutionName struct {
			Auto    string `json:"AUTO"`
			FullHd1 string `json:"FULL_HD1"`
			Hd1     string `json:"HD1"`
			Origion string `json:"ORIGION"`
			Sd1     string `json:"SD1"`
			Sd2     string `json:"SD2"`
		} `json:"resolution_name"`
		RtmpPullURL       string  `json:"rtmp_pull_url"`
		RtmpPullURLParams string  `json:"rtmp_pull_url_params"`
		RtmpPushURL       string  `json:"rtmp_push_url"`
		RtmpPushURLParams string  `json:"rtmp_push_url_params"`
		StreamControlType float64 `json:"stream_control_type"`
	} `json:"stream_url"`
	StreamURLFilteredInfo struct {
		IsGatedRoom bool `json:"is_gated_room"`
		IsPaidEvent bool `json:"is_paid_event"`
	} `json:"stream_url_filtered_info"`
	Title             string    `json:"title"`
	TopFans           []*TopFan `json:"top_fans"`
	UseFilter         bool      `json:"use_filter"`
	UserCount         float64   `json:"user_count"`
	UserShareText     string    `json:"user_share_text"`
	VideoFeedTag      string    `json:"video_feed_tag"`
	WebcastCommentTcs float64   `json:"webcast_comment_tcs"`
	WebcastSdkVersion float64   `json:"webcast_sdk_version"`
	WithDrawSomething bool      `json:"with_draw_something"`
	WithKtv           bool      `json:"with_ktv"`
	WithLinkmic       bool      `json:"with_linkmic"`
}

type GiftInfoRsp struct {
	GiftInfo *GiftInfo `json:"data"`
	Extra    struct {
		LogID string `json:"log_id"`
		Now   int64  `json:"now"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

type GiftInfo struct {
	DoodleTemplates []interface{} `json:"doodle_templates"`
	Gifts           []struct {
		ActionType            int           `json:"action_type"`
		AppID                 int           `json:"app_id"`
		BusinessText          string        `json:"business_text"`
		ColorInfos            []interface{} `json:"color_infos"`
		Combo                 bool          `json:"combo"`
		Describe              string        `json:"describe"`
		DiamondCount          int           `json:"diamond_count"`
		Duration              int           `json:"duration"`
		EventName             string        `json:"event_name"`
		ForCustom             bool          `json:"for_custom"`
		ForLinkmic            bool          `json:"for_linkmic"`
		GiftRankRecommendInfo string        `json:"gift_rank_recommend_info"`
		GiftScene             int           `json:"gift_scene"`
		GoldEffect            string        `json:"gold_effect"`
		GraySchemeURL         string        `json:"gray_scheme_url"`
		GuideURL              string        `json:"guide_url"`
		Icon                  struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"icon"`
		ID    int `json:"id"`
		Image struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"image"`
		IsBroadcastGift    bool `json:"is_broadcast_gift"`
		IsDisplayedOnPanel bool `json:"is_displayed_on_panel"`
		IsEffectBefview    bool `json:"is_effect_befview"`
		IsGray             bool `json:"is_gray"`
		IsRandomGift       bool `json:"is_random_gift"`
		ItemType           int  `json:"item_type"`
		LockInfo           struct {
			Lock     bool `json:"lock"`
			LockType int  `json:"lock_type"`
		} `json:"lock_info"`
		Manual          string `json:"manual"`
		Name            string `json:"name"`
		Notify          bool   `json:"notify"`
		PrimaryEffectID int    `json:"primary_effect_id"`
		Region          string `json:"region"`
		SchemeURL       string `json:"scheme_url"`
		SpecialEffects  struct {
		} `json:"special_effects"`
		TriggerWords  []interface{} `json:"trigger_words"`
		Type          int           `json:"type"`
		GiftLabelIcon struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"gift_label_icon,omitempty"`
		PreviewImage struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"preview_image,omitempty"`
		TrackerParams struct {
			GiftProperty string `json:"gift_property"`
		} `json:"tracker_params,omitempty"`
		GiftPanelBanner struct {
			BgColorValues []interface{} `json:"bg_color_values"`
			DisplayText   struct {
				DefaultFormat struct {
					Bold               bool   `json:"bold"`
					Color              string `json:"color"`
					FontSize           int    `json:"font_size"`
					Italic             bool   `json:"italic"`
					ItalicAngle        int    `json:"italic_angle"`
					UseHeighLightColor bool   `json:"use_heigh_light_color"`
					UseRemoteClor      bool   `json:"use_remote_clor"`
					Weight             int    `json:"weight"`
				} `json:"default_format"`
				DefaultPattern string        `json:"default_pattern"`
				Key            string        `json:"key"`
				Pieces         []interface{} `json:"pieces"`
			} `json:"display_text"`
			LeftIcon struct {
				AvgColor   string   `json:"avg_color"`
				Height     int      `json:"height"`
				ImageType  int      `json:"image_type"`
				IsAnimated bool     `json:"is_animated"`
				OpenWebURL string   `json:"open_web_url"`
				URI        string   `json:"uri"`
				URLList    []string `json:"url_list"`
				Width      int      `json:"width"`
			} `json:"left_icon"`
			SchemaURL string `json:"schema_url"`
		} `json:"gift_panel_banner,omitempty"`
	} `json:"gifts"`
	GiftsInfo struct {
		ColorGiftIconAnimation struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"color_gift_icon_animation"`
		DefaultLocColorGiftID            int  `json:"default_loc_color_gift_id"`
		EnableFirstRechargeDynamicEffect bool `json:"enable_first_recharge_dynamic_effect"`
		FirstRechargeGiftInfo            struct {
			ExpireAt             int `json:"expire_at"`
			GiftID               int `json:"gift_id"`
			OriginalDiamondCount int `json:"original_diamond_count"`
		} `json:"first_recharge_gift_info"`
		GiftComboInfos []interface{} `json:"gift_combo_infos"`
		GiftGroupInfos []struct {
			GroupCount int    `json:"group_count"`
			GroupText  string `json:"group_text"`
		} `json:"gift_group_infos"`
		GiftIconInfo struct {
			EffectURI string `json:"effect_uri"`
			Icon      struct {
				AvgColor   string        `json:"avg_color"`
				Height     int           `json:"height"`
				ImageType  int           `json:"image_type"`
				IsAnimated bool          `json:"is_animated"`
				OpenWebURL string        `json:"open_web_url"`
				URI        string        `json:"uri"`
				URLList    []interface{} `json:"url_list"`
				Width      int           `json:"width"`
			} `json:"icon"`
			IconID       int    `json:"icon_id"`
			IconURI      string `json:"icon_uri"`
			Name         string `json:"name"`
			ValidEndAt   int    `json:"valid_end_at"`
			ValidStartAt int    `json:"valid_start_at"`
			WithEffect   bool   `json:"with_effect"`
		} `json:"gift_icon_info"`
		GiftPollInfo struct {
			GiftPollOptions []struct {
				GiftID         int `json:"gift_id"`
				PollResultIcon struct {
					AvgColor   string   `json:"avg_color"`
					Height     int      `json:"height"`
					ImageType  int      `json:"image_type"`
					IsAnimated bool     `json:"is_animated"`
					OpenWebURL string   `json:"open_web_url"`
					URI        string   `json:"uri"`
					URLList    []string `json:"url_list"`
					Width      int      `json:"width"`
				} `json:"poll_result_icon"`
			} `json:"gift_poll_options"`
		} `json:"gift_poll_info"`
		GiftWords                 string `json:"gift_words"`
		HideRechargeEntry         bool   `json:"hide_recharge_entry"`
		NewGiftID                 int    `json:"new_gift_id"`
		RecentlySentColorGiftID   int    `json:"recently_sent_color_gift_id"`
		RecommendedRandomGiftID   int    `json:"recommended_random_gift_id"`
		ShowFirstRechargeEntrance bool   `json:"show_first_recharge_entrance"`
		SpeedyGiftID              int    `json:"speedy_gift_id"`
	} `json:"gifts_info"`
	Pages []interface{} `json:"pages"`
}

type TopFan struct {
	FanTicket float64 `json:"fan_ticket"`
	User      struct {
		AllowFindByContacts                 bool `json:"allow_find_by_contacts"`
		AllowOthersDownloadVideo            bool `json:"allow_others_download_video"`
		AllowOthersDownloadWhenSharingVideo bool `json:"allow_others_download_when_sharing_video"`
		AllowShareShowProfile               bool `json:"allow_share_show_profile"`
		AllowShowInGossip                   bool `json:"allow_show_in_gossip"`
		AllowShowMyAction                   bool `json:"allow_show_my_action"`
		AllowStrangeComment                 bool `json:"allow_strange_comment"`
		AllowUnfollowerComment              bool `json:"allow_unfollower_comment"`
		AllowUseLinkmic                     bool `json:"allow_use_linkmic"`
		AvatarLarge                         struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"avatar_large"`
		AvatarMedium struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"avatar_medium"`
		AvatarThumb struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"avatar_thumb"`
		BadgeImageList           []interface{} `json:"badge_image_list"`
		BadgeList                []interface{} `json:"badge_list"`
		BgImgURL                 string        `json:"bg_img_url"`
		BioDescription           string        `json:"bio_description"`
		BlockStatus              float64       `json:"block_status"`
		BorderList               []interface{} `json:"border_list"`
		CommentRestrict          float64       `json:"comment_restrict"`
		CommerceWebcastConfigIds []interface{} `json:"commerce_webcast_config_ids"`
		Constellation            string        `json:"constellation"`
		CreateTime               float64       `json:"create_time"`
		DisableIchat             float64       `json:"disable_ichat"`
		DisplayID                string        `json:"display_id"`
		EnableIchatImg           float64       `json:"enable_ichat_img"`
		Exp                      float64       `json:"exp"`
		FanTicketCount           float64       `json:"fan_ticket_count"`
		FoldStrangerChat         bool          `json:"fold_stranger_chat"`
		FollowInfo               struct {
			FollowStatus   float64 `json:"follow_status"`
			FollowerCount  float64 `json:"follower_count"`
			FollowingCount float64 `json:"following_count"`
			PushStatus     float64 `json:"push_status"`
		} `json:"follow_info"`
		FollowStatus        float64       `json:"follow_status"`
		IchatRestrictType   float64       `json:"ichat_restrict_type"`
		ID                  int64         `json:"id"`
		IDStr               string        `json:"id_str"`
		IsFollower          bool          `json:"is_follower"`
		IsFollowing         bool          `json:"is_following"`
		LinkMicStats        float64       `json:"link_mic_stats"`
		MediaBadgeImageList []interface{} `json:"media_badge_image_list"`
		ModifyTime          float64       `json:"modify_time"`
		NeedProfileGuide    bool          `json:"need_profile_guide"`
		NewRealTimeIcons    []interface{} `json:"new_real_time_icons"`
		Nickname            string        `json:"nickname"`
		PayGrade            struct {
			GradeBanner        string        `json:"grade_banner"`
			GradeDescribe      string        `json:"grade_describe"`
			GradeIconList      []interface{} `json:"grade_icon_list"`
			Level              float64       `json:"level"`
			Name               string        `json:"name"`
			NextName           string        `json:"next_name"`
			NextPrivileges     string        `json:"next_privileges"`
			Score              float64       `json:"score"`
			ScreenChatType     float64       `json:"screen_chat_type"`
			UpgradeNeedConsume float64       `json:"upgrade_need_consume"`
		} `json:"pay_grade"`
		PayScore           float64       `json:"pay_score"`
		PayScores          float64       `json:"pay_scores"`
		PushCommentStatus  bool          `json:"push_comment_status"`
		PushDigg           bool          `json:"push_digg"`
		PushFollow         bool          `json:"push_follow"`
		PushFriendAction   bool          `json:"push_friend_action"`
		PushIchat          bool          `json:"push_ichat"`
		PushStatus         bool          `json:"push_status"`
		PushVideoPost      bool          `json:"push_video_post"`
		PushVideoRecommend bool          `json:"push_video_recommend"`
		RealTimeIcons      []interface{} `json:"real_time_icons"`
		SecUID             string        `json:"sec_uid"`
		Secret             float64       `json:"secret"`
		ShareQrcodeURI     string        `json:"share_qrcode_uri"`
		SpecialID          string        `json:"special_id"`
		Status             float64       `json:"status"`
		TicketCount        float64       `json:"ticket_count"`
		TopFans            []*TopFan     `json:"top_fans"`
		TopVipNo           float64       `json:"top_vip_no"`
		UserAttr           struct {
			IsAdmin      bool    `json:"is_admin"`
			IsMuted      bool    `json:"is_muted"`
			IsSuperAdmin bool    `json:"is_super_admin"`
			MuteDuration float64 `json:"mute_duration"`
		} `json:"user_attr"`
		UserRole                    float64 `json:"user_role"`
		Verified                    bool    `json:"verified"`
		VerifiedContent             string  `json:"verified_content"`
		VerifiedReason              string  `json:"verified_reason"`
		WithCarManagementPermission bool    `json:"with_car_management_permission"`
		WithCommercePermission      bool    `json:"with_commerce_permission"`
		WithFusionShopEntry         bool    `json:"with_fusion_shop_entry"`
	} `json:"user"`
}
