package gotiktoklive

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	pb "gotiktoklive/proto"

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

	cookies := l.tiktok.c.Jar.Cookies(u)
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

	url := fmt.Sprintf("%s?%s", addr, vs.Encode())
	dialer := ws.Dialer{
		Header: ws.HandshakeHeaderHTTP(headers),
		NetDial: func(ctx context.Context, a, b string) (net.Conn, error) {
			return proxy.FromEnvironment().Dial(a, b)
		},
		Protocols: []string{"echo-protocol"},
	}
	conn, _, _, err := dialer.Dial(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Failed to connect: %w", err)
	}
	l.wss = conn
	return nil
}

func (l *Live) readSocket() {
	defer l.wss.Close()

	want := ws.OpBinary
	s := ws.StateClientSide

	// TODO: Tidy up duplication
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
			handleReadError(err)
		}

		// If msg is ping or close
		if hdr.OpCode.IsControl() {
			if err := controlHandler(hdr, &rd); err != nil {
				handleReadError(err)
			}
			continue
		}

		// Reopen connection if it was closed
		if hdr.OpCode == ws.OpClose {
			panic("Socket closed")
		}

		// Wrong OpCode
		if hdr.OpCode&want == 0 {
			if err := rd.Discard(); err != nil {
				handleReadError(err)
			}
			continue
		}

		// Read message
		msgBytes, err := ioutil.ReadAll(&rd)
		if err != nil {
			handleReadError(err)
		}

		if err := l.parseWssMsg(msgBytes); err != nil {
			handleReadError(err)
		}

		// Gracefully shutdown
		// if f.done == true {
		// 	return
		// }
	}
}

func handleReadError(err error) {
	panic(err)
}

func (l *Live) parseWssMsg(wssMsg []byte) error {
	var rsp pb.WebcastWebsocketMessage
	if err := proto.Unmarshal(wssMsg, &rsp); err != nil {
		return err
	}

	if rsp.Type == "msg" {
		var response pb.WebcastResponse
		if err := proto.Unmarshal(rsp.Binary, &response); err != nil {
			return err
		}

		if err := l.sendAck(rsp.Id); err != nil {
			return err
		}
		l.cursor = response.Cursor

		if l.tiktok.Debug {
			l.tiktok.debugHandler(fmt.Sprintf("Got %d messages, %s", len(response.Messages), response.Cursor))
		}

		for _, rawMsg := range response.Messages {
			msg, err := parseMsg(rawMsg, l.tiktok.warnHandler)
			if err != nil {
				return err
			}
			if msg != nil {
				l.Events <- msg
			}
		}
		return nil
	}
	if l.tiktok.Debug {
		l.tiktok.debugHandler(fmt.Sprintf("Message type unknown, %s : '%s'", rsp.Type, string(rsp.Binary)))
	}
	return nil
}

func (l *Live) sendPing() {
	b, err := hex.DecodeString("3A026862")
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(10000 * time.Millisecond)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if err := wsutil.WriteClientBinary(l.wss, b); err != nil {
				panic(err)
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
