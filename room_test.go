package gotiktoklive

import (
	"testing"
)

func TestRoomID(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID("jbvcreative")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found Room ID: %d", id)
}
