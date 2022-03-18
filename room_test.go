package gotiktoklive

import (
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
