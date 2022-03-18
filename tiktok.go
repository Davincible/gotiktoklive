package gotiktoklive

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func (t *TikTok) getRoomID(user string) (int64, error) {
	body, err := t.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf(urlUserLive, user),
		OmitAPI:  true,
	})
	if err != nil {
		return 0, err
	}

	if id := reRoomIDMeta.FindSubmatch(body); len(id) > 1 {
		return strconv.ParseInt(string(id[1]), 10, 64)
	}

	if id := reRoomIDJson.FindSubmatch(body); len(id) > 1 {
		return strconv.ParseInt(string(id[1]), 10, 64)
	}

	if strings.Contains(string(body), `"og:url"`) {
		return 0, ErrUserOffline
	}
	return 0, ErrIPBlocked
}

func defaultHandler(i ...interface{}) {
	fmt.Println(i...)
}
