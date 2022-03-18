package gotiktoklive

import (
	"encoding/base64"
	"testing"
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

	info, err := tiktok.getRoomInfo(id)
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

	info, err := tiktok.getGiftInfo(id)
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

	info, err := tiktok.getRoomData(id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Ws url: %s, %+v", info.WsUrl, info.WsParam)

	for _, msg := range info.Messages {
		if msg.Type == "WebcastLiveIntroMessage" {
			m := base64.StdEncoding.EncodeToString(msg.Binary)
			t.Logf("Message Type: %s : %v", msg.Type, m)
		}
	}
}
