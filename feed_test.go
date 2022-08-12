package gotiktoklive

import (
	"sort"
	"testing"

	"github.com/Davincible/gotiktoklive/tests"
)

func TestFeedItem(t *testing.T) {
	tiktok := NewTikTok()
	feed := tiktok.NewFeed()

	items := []*LiveStream{}
	i := 0
	for {
		feedItem, err := feed.Next()
		if err != nil {
			t.Fatal(err)
		}
		items = append(items, feedItem.LiveStreams...)
		i++
		t.Logf("%d : %d, %v", feedItem.Extra.MaxTime, len(feedItem.LiveStreams), feedItem.Extra.HasMore)
		for _, stream := range feedItem.LiveStreams {
			t.Logf("%s : %d viewers, %s", stream.Room.Owner.Nickname, stream.Room.UserCount, stream.LiveReason)
		}

		if !feedItem.Extra.HasMore || i > 5 {
			break
		}
	}
	t.Logf("Found %d items, over %d requests", len(items), i)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Room.UserCount > items[j].Room.UserCount
	})

	tests.USERNAME = items[0].Room.Owner.Username
	t.Logf("Setting username to %s", items[0].Room.Owner.Username)
}
