package gotiktoklive

import (
	"testing"
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
		t:  tiktok,
		ID: id,
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
		t:  tiktok,
		ID: id,
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
		t:      tiktok,
		ID:     id,
		Events: make(chan interface{}, 100),
	}

	err = live.getRoomData()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Ws url: %s, %+v", live.wsURL, live.wsParams)
}
