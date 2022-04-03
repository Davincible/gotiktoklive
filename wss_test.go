package gotiktoklive

import (
	"testing"
	"time"
)

func TestWebsocket(t *testing.T) {
	tiktok := NewTikTok()
	tiktok.Debug = true
	tiktok.debugHandler = func(i ...interface{}) {
		t.Log(i...)
	}
	id, err := tiktok.getRoomID(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	live := Live{
		t:      tiktok,
		ID:     id,
		Events: make(chan interface{}, 100),
	}

	err = live.getRoomData()
	if err != nil {
		t.Fatal(err)
	}

	if live.wsURL == "" {
		t.Fatal("No websocket url provided")
	}
	t.Logf("Ws url: %s, %+v", live.wsURL, live.wsParams)

	if err := live.connect(live.wsURL, live.wsParams); err != nil {
		t.Fatal(err)
	}

	tiktok.wg.Add(2)
	go live.readSocket()
	go live.sendPing()

	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			return
		case event := <-live.Events:
			t.Logf("%T: %+v", event, event)
		}
	}
}
