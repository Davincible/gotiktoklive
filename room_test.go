package gotiktoklive

import (
	"encoding/json"
	"testing"
	"time"
)

const (
	// USERNAME = "jbvcreative"
	USERNAME = "promobot.robots"
)

func TestRoomID(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found Room ID: %s ", id)
}

func TestRoomInfo(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	live := Live{
		tiktok: tiktok,
		ID:     id,
	}

	info, err := live.getRoomInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(info.Title)
}

func TestGiftInfo(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	live := Live{
		tiktok: tiktok,
		ID:     id,
	}

	info, err := live.getGiftInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found %d gifts", len(info.Gifts))
}

func TestRoomData(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	live := Live{
		tiktok: tiktok,
		ID:     id,
	}

	info, err := live.getRoomData()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Ws url: %s, %+v", info.WsUrl, info.WsParam)
}

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
		tiktok: tiktok,
		ID:     id,
		Events: make(chan interface{}, 100),
	}

	info, err := live.getRoomData()
	if err != nil {
		t.Fatal(err)
	}

	if info.WsUrl == "" {
		t.Fatal("No websocket url provided")
	}
	t.Logf("Ws url: %s, %+v", info.WsUrl, info.WsParam)

	if err := live.connect(info.WsUrl, map[string]string{info.WsParam.Name: info.WsParam.Value}); err != nil {
		t.Fatal(err)
	}
	go live.readSocket()
	go live.sendPing()

	timeout := time.After(10 * time.Second)
	for {
		select {
		case <-timeout:
			return
		case event := <-live.Events:
			b, err := json.Marshal(event)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("%T: %s", event, string(b))
		}
	}

}
