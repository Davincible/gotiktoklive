package gotiktoklive

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	pb "github.com/Davincible/gotiktoklive/proto"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"golang.org/x/net/proxy"
	"google.golang.org/protobuf/proto"
)

func (l *Live) connect(addr string, params map[string]string) error {
	u, err := url.Parse("https://tiktok.com/")
	if err != nil {
		return nil
	}

	// Take the cookies from the HTTP client cookie jar and pass them as a GET parameter.
	cookies := l.t.c.Jar.Cookies(u)
	headers := http.Header{}
	var s string
	for _, cookie := range cookies {
		if s == "" {
			s = cookie.String()
		} else {
			s += "; " + cookie.String()
		}
	}
	if len(cookies) > 0 {
		headers.Set("Cookie", s)
	}

	vs := url.Values{}
	for key, value := range defaultGETParams {
		vs.Add(key, value)
	}
	for key, value := range params {
		vs.Add(key, value)
	}
	vs.Add("room_id", l.ID)

	wsURL := fmt.Sprintf("%s?%s", addr, vs.Encode())

	// Read proxies from HTTP env variables, HTTPS takes precedent
	var proxyURI *url.URL
	envVars := []string{"HTTP_PROXY", "HTTPS_PROXY"}
	for _, envVar := range envVars {
		if httpProxy := os.Getenv(envVar); httpProxy != "" {
			uri, err := url.Parse(httpProxy)
			if err != nil {
				l.t.warnHandler(fmt.Errorf("Failed to parse %s url: %w", envVar, err))
			} else {
				proxyURI = uri
			}
		}
	}

	// If proxy manually defined, use that
	if l.t.proxy != nil {
		proxyURI = l.t.proxy
	}

	// Default proxy net dial
	proxyNetDial := proxy.FromEnvironment()

	// If custom URI is defined, genereate proxy net dial from that
	if proxyURI != nil {
		dial, err := proxy.FromURL(proxyURI, Direct)
		if err != nil {
			l.t.warnHandler(fmt.Errorf("Failed to configure proxy dialer: %w", err))
		} else {
			proxyNetDial = dial
		}
	}

	dialer := ws.Dialer{
		Header: ws.HandshakeHeaderHTTP(headers),
		NetDial: func(ctx context.Context, a, b string) (net.Conn, error) {
			return proxyNetDial.Dial(a, b)
		},
		// NetDial:   proxy.Dial,
		Protocols: []string{"echo-protocol"},
	}
	conn, _, _, err := dialer.Dial(context.Background(), wsURL)
	if err != nil {
		return fmt.Errorf("Failed to connect: %w", err)
	}
	l.wss = conn
	return nil
}

func (l *Live) readSocket() {
	defer l.wss.Close()
	defer l.t.wg.Done()

	want := ws.OpBinary
	s := ws.StateClientSide

	controlHandler := wsutil.ControlFrameHandler(l.wss, s)
	rd := wsutil.Reader{
		Source:          l.wss,
		State:           s,
		CheckUTF8:       true,
		SkipHeaderCheck: false,
		OnIntermediate:  controlHandler,
	}

	for {
		hdr, err := rd.NextFrame()
		if err != nil {
			l.t.errHandler(fmt.Errorf("Failed to read websocket message"))
		}

		// If msg is ping or close
		if hdr.OpCode.IsControl() {
			if err := controlHandler(hdr, &rd); err != nil {
				l.t.errHandler(err)
			}
			continue
		}

		// Reopen connection if it was closed
		if hdr.OpCode == ws.OpClose {
			connected, err := l.tryConnectionUpgrade()
			if err != nil {
				l.t.errHandler(err)
			}
			if !connected {
				l.t.wg.Add(1)
				go l.startPolling()
				return
			}
		}

		// Wrong OpCode
		if hdr.OpCode&want == 0 {
			if err := rd.Discard(); err != nil {
				l.t.errHandler(err)
			}
			continue
		}

		// Read message
		msgBytes, err := ioutil.ReadAll(&rd)
		if err != nil {
			l.t.errHandler(err)
		}

		if err := l.parseWssMsg(msgBytes); err != nil {
			l.t.errHandler(err)
		}

		// Gracefully shutdown
		select {
		case <-l.done():
			return
		case <-l.t.done():
			return
		default:
		}
	}
}

func (l *Live) parseWssMsg(wssMsg []byte) error {
	var rsp pb.WebcastWebsocketMessage
	if err := proto.Unmarshal(wssMsg, &rsp); err != nil {
		return fmt.Errorf("Failed to unmarshal proto WebcastWebsocketMessage: %w", err)
	}

	if rsp.Type == "msg" {
		var response pb.WebcastResponse
		if err := proto.Unmarshal(rsp.Binary, &response); err != nil {
			return fmt.Errorf("Failed to unmarshal proto WebcastResponse: %w", err)
		}

		if err := l.sendAck(rsp.Id); err != nil {
			return err
		}
		l.cursor = response.Cursor

		if l.t.Debug {
			l.t.debugHandler(fmt.Sprintf("Got %d messages, %s", len(response.Messages), response.Cursor))
		}

		for _, rawMsg := range response.Messages {
			msg, err := parseMsg(rawMsg, l.t.warnHandler)
			if err != nil {
				return err
			}
			if msg != nil {
				l.Events <- msg
			}

			// If livestream has ended
			if m, ok := msg.(ControlEvent); ok && m.Action == 3 {
				go func() {
					select {
					case <-time.After(10 * time.Second):
						l.close()
					}
				}()
			}
		}
		return nil
	}
	if l.t.Debug {
		l.t.debugHandler(fmt.Sprintf("Message type unknown, %s : '%s'", rsp.Type, string(rsp.Binary)))
	}
	return nil
}

func (l *Live) sendPing() {
	defer l.t.wg.Done()

	b, err := hex.DecodeString("3A026862")
	if err != nil {
		l.t.errHandler(err)
	}

	t := time.NewTicker(10000 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-l.done():
			return
		case <-l.t.done():
			return
		case <-t.C:
			if err := wsutil.WriteClientBinary(l.wss, b); err != nil {
				l.t.errHandler(err)
			}
		}
	}
}

func (l *Live) sendAck(id uint64) error {
	msg := pb.WebcastWebsocketAck{
		Id:   id,
		Type: "ack",
	}

	b, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}

	if err := wsutil.WriteClientBinary(l.wss, b); err != nil {
		return err
	}
	return nil
}

func (l *Live) tryConnectionUpgrade() (bool, error) {
	if l.wsURL != "" && l.wsParams != nil {
		err := l.connect(l.wsURL, l.wsParams)
		if err != nil {
			return false, err
		}
		if l.t.Debug {
			l.t.debugHandler("Connected to websocket")
		}

		l.t.wg.Add(2)
		go l.readSocket()
		go l.sendPing()

		return true, nil
	}
	return false, nil
}
