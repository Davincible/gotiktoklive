package gotiktoklive

import (
	"errors"
	"regexp"
)

const (
	// Base URL
	tiktokBaseUrl = "https://www.tiktok.com/"
	tiktokAPIUrl  = "https://webcast.tiktok.com/webcast/"
	tiktokSigner  = "https://tiktok.isaackogan.com/"

	// Endpoints
	urlLive      = "live/"
	urlFeed      = "feed/"
	urlRankList  = "ranklist/online_audience/"
	urlPriceList = "diamond/"
	urlUser      = "@%s/"
	// Think this changed to room/enter/
	urlRoomInfo = "room/info/"
	urlRoomData = "im/fetch/"
	urlGiftInfo = "gift/list/"
	urlSignReq  = "webcast/sign_url/"

	// Default Request Headers
	userAgent = "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36"
	referer   = "https://www.tiktok.com/"
	origin    = "https://www.tiktok.com"
	clientId  = "ttlive-golang"
)

var (
	// Default GET Request parameters
	defaultGETParams = map[string]string{
		"aid":                 "1988",
		"app_language":        "en-US",
		"app_name":            "tiktok_web",
		"browser_language":    "en",
		"browser_name":        "Mozilla",
		"browser_online":      "true",
		"browser_platform":    "Win32",
		"browser_version":     "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36",
		"cookie_enabled":      "true",
		"cursor":              "",
		"internal_ext":        "",
		"device_platform":     "web",
		"focus_state":         "true",
		"from_page":           "user",
		"history_len":         "4",
		"is_fullscreen":       "false",
		"is_page_visible":     "true",
		"did_rule":            "3",
		"fetch_rule":          "1",
		"identity":            "audience",
		"last_rtt":            "0",
		"live_id":             "12",
		"resp_content_type":   "protobuf",
		"screen_height":       "1152",
		"screen_width":        "2048",
		"tz_name":             "Europe/Berlin",
		"referer":             "https://www.tiktok.com/",
		"root_referer":        "https://www.tiktok.com/",
		"version_code":        "180800",
		"webcast_sdk_version": "1.3.0",
		"update_version_code": "1.3.0",
	}
	reJsonData = []*regexp.Regexp{
		regexp.MustCompile(`<script id="SIGI_STATE"[^>]+>(.*?)</script>`),
		regexp.MustCompile(`<script id="sigi-persisted-data">window\['SIGI_STATE'\]=(.*);w`),
	}
	reVerify = regexp.MustCompile(`tiktok-verify-page`)
)

var (
	ErrUserOffline       = errors.New("User might be offline, Room ID not found")
	ErrIPBlocked         = errors.New("Your IP or country might be blocked by TikTok.")
	ErrLiveHasEnded      = errors.New("Livestream has ended")
	ErrMsgNotImplemented = errors.New("Message protobuf type has not been implemented, please report")
	ErrNoMoreFeedItems   = errors.New("No more feed items available")
	ErrUserNotFound      = errors.New("User not found")
	ErrCaptcha           = errors.New("Captcha detected, unable to proceed")
	ErrUrlNotFound       = errors.New("Unable to download stream, URL not found.")
	ErrFFMPEGNotFound    = errors.New("Please install ffmpeg before downloading.")
	ErrRateLimitExceeded = errors.New("You have exceeded the rate limit, please wait a few min")
)
