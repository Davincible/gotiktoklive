package gotiktoklive

import (
	"errors"
	"regexp"
)

const (
	// Base URL
	tiktokBaseUrl = "https://www.tiktok.com/"
	tiktokAPIUrl  = "https://webcast.tiktok.com/webcast/"

	// Endpoints
	urlLive      = "live/"
	urlFeed      = "feed/"
	urlRankList  = "ranklist/online_audience/"
	urlPriceList = "diamond/"
	urlUser      = "@%s/"
	urlRoomInfo  = "room/info/"
	urlRoomData  = "im/fetch/"
	urlGiftInfo  = "gift/list/"

	// Default Request Headers
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"
	referer   = "https://www.tiktok.com/"
	origin    = "https://www.tiktok.com"
)

var (
	// Default GET Request parameters
	defaultGETParams = map[string]string{
		"aid":               "1988",
		"app_language":      "en-US",
		"app_name":          "tiktok_web",
		"browser_language":  "en",
		"browser_name":      "Mozilla",
		"browser_online":    "true",
		"browser_platform":  "Win32",
		"browser_version":   "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36",
		"cookie_enabled":    "true",
		"cursor":            "",
		"device_platform":   "web",
		"did_rule":          "3",
		"fetch_rule":        "1",
		"identity":          "audience",
		"internal_ext":      "",
		"last_rtt":          "0",
		"live_id":           "12",
		"resp_content_type": "protobuf",
		"screen_height":     "1152",
		"screen_width":      "2048",
		"tz_name":           "Europe/Berlin",
		"version_code":      "180800",
	}

	reRoomIDMeta = regexp.MustCompile("room_id=([0-9]*)")
	reRoomIDJson = regexp.MustCompile(`"roomId":"([0-9]*)"`)
	reJsonData   = regexp.MustCompile(`<script id="sigi-persisted-data">window\['SIGI_STATE'\]=(.*);w`)
)

var (
	ErrUserOffline       = errors.New("User might be offline, Room ID not found")
	ErrIPBlocked         = errors.New("Your IP or country might be blocked by TikTok.")
	ErrLiveHasEnded      = errors.New("Livestream has ended")
	ErrMsgNotImplemented = errors.New("Message protobuf type has not been implemented, please report")
	ErrNoMoreFeedItems   = errors.New("No more feed items available")
	ErrUserNotFound      = errors.New("User not found")
)
