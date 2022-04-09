package gotiktoklive

import (
	"encoding/json"
	"strconv"
)

// Feed allows you to fetch reccomended livestreams.
type Feed struct {
	t *TikTok

	// All collected reccomended livestreams
	LiveStreams []*LiveStream

	HasMore bool
	maxTime int64
}

// NewFeed creates a new Feed instance. Start fetching reccomended livestreams
//  with Feed.Next().
func (t *TikTok) NewFeed() *Feed {
	return &Feed{
		t:           t,
		LiveStreams: []*LiveStream{},
		HasMore:     true,
	}
}

// Next fetches the next couple of recommended live streams, if available.
// You can call this as long as Feed.HasMore = true. All items will be added
//  to the Feed.LiveStreams list.
func (f *Feed) Next() (*FeedItem, error) {
	if !f.HasMore {
		return nil, ErrNoMoreFeedItems
	}

	params := copyMap(defaultGETParams)
	params["channel"] = "tiktok_web"
	params["channel_id"] = "86"
	if f.maxTime != 0 {
		params["max_time"] = strconv.FormatInt(f.maxTime, 10)
	}

	body, err := f.t.sendRequest(&reqOptions{
		Endpoint: urlFeed,
		Query:    params,
	})
	if err != nil {
		return nil, err
	}

	var rsp FeedItem
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}

	f.HasMore = rsp.Extra.HasMore
	f.maxTime = rsp.Extra.MaxTime
	for _, s := range rsp.LiveStreams {
		s.t = f.t
	}

	return &rsp, nil
}

// Track stars tracking the livestream obtained from the Feed, and returns
//  a Live instance, just as if you would start tracking the user with
//  tiktok.TrackUser(<user>).
func (s *LiveStream) Track() (*Live, error) {
	return s.t.TrackRoom(s.Rid)
}
