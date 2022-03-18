package gotiktoklive

import (
	"testing"
)

const (
	// USERNAME = "jbvcreative"
	USERNAME = "thienlongpro87"
)

func TestRoomID(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found Room ID: %s", id)
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
