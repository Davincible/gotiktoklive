package tests

import (
	"testing"

	"github.com/Davincible/gotiktoklive"
)

func TestUserInfo(t *testing.T) {
	tiktok := gotiktoklive.NewTikTok()

	info, err := tiktok.GetUserInfo(USERNAME)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Got user info: %+v", info)
}
