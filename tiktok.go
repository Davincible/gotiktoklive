package gotiktoklive

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// TikTok allows you to track and discover current live streams.
type TikTok struct {
	c    *http.Client
	wg   *sync.WaitGroup
	done func() <-chan struct{}

	// Pass extra debug messages to debugHandler
	Debug bool

	// LogRequests when set to true will log all made requests in JSON to debugHandler
	LogRequests bool

	infoHandler  func(...interface{})
	warnHandler  func(...interface{})
	debugHandler func(...interface{})

	errHandler func(error)
}

// NewTikTok creates a tiktok instance that allows you to track live streams and
//  discover current livestreams.
func NewTikTok() *TikTok {
	jar, _ := cookiejar.New(nil)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tiktok := TikTok{
		c:            &http.Client{Jar: jar},
		wg:           &wg,
		done:         ctx.Done,
		infoHandler:  defaultLogHandler,
		warnHandler:  defaultLogHandler,
		debugHandler: defaultLogHandler,
		errHandler:   routineErrHandler,
	}

	setupInterruptHandler(
		func(c chan os.Signal) {
			<-c
			cancel()
			wg.Wait()

			tiktok.infoHandler("Shutting down...")
			os.Exit(0)
		})

	return &tiktok
}

// GetPriceList fetches the price list of tiktok coins. Prices will be given in
//  USD cents and the cents equivalent of the local currency of the IP location.
// To fetch a different currency, use a VPN or proxy to change your IP to a
//  different country.
func (t *TikTok) GetPriceList() (*PriceList, error) {
	body, err := t.sendRequest(&reqOptions{
		Endpoint: urlPriceList,
		Query:    defaultGETParams,
	})
	if err != nil {
		return nil, err
	}

	var rsp PriceList
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

func (t *TikTok) SetInfoHandler(f func(...interface{})) {
	t.infoHandler = f
}

func (t *TikTok) SetWarnHandler(f func(...interface{})) {
	t.warnHandler = f
}

func (t *TikTok) SetDebugHandler(f func(...interface{})) {
	t.debugHandler = f
}

func (t *TikTok) SetErrorHandler(f func(error)) {
	t.errHandler = f
}

func setupInterruptHandler(f func(chan os.Signal)) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go f(c)
}
