package gotiktoklive

import (
	"testing"

	"github.com/Davincible/gotiktoklive/tests"
)

func TestRoomID(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(tests.USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Found Room ID: %s ", id)
}

func TestRoomInfo(t *testing.T) {
	tiktok := NewTikTok()
	id, err := tiktok.getRoomID(tests.USERNAME)
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
	id, err := tiktok.getRoomID(tests.USERNAME)
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
	id, err := tiktok.getRoomID(tests.USERNAME)
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

// func TestHeadless(t *testing.T) {
// 	tiktok := NewTikTok()
// 	tiktok.SetProxy("http://127.0.0.1:8080", false)
// 	// err := tiktok.openTikTok("https://intoli.com/blog/not-possible-to-block-chrome-headless/chrome-headless-test.html")
// 	err := tiktok.openTikTok("https://www.tiktok.com/")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	t.Log("Test done!")
// }
