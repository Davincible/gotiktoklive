package gotiktoklive

import (
	"net/http"
	"net/http/cookiejar"
)

type TikTok struct {
	c *http.Client

	Debug bool

	infoHandler  func(...interface{})
	warnHandler  func(...interface{})
	debugHandler func(...interface{})
}

func NewTikTok() *TikTok {
	jar, _ := cookiejar.New(nil)

	return &TikTok{
		c:            &http.Client{Jar: jar},
		infoHandler:  defaultHandler,
		warnHandler:  defaultHandler,
		debugHandler: defaultHandler,
	}
}
