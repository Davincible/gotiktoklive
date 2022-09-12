[![GoDoc](https://godoc.org/github.com/Davincible/gotiktoklive?status.svg)](https://godoc.org/github.com/Davincible/gotiktoklive) [![Go Report Card](https://goreportcard.com/badge/github.com/Davincible/gotiktoklive)](https://goreportcard.com/report/github.com/Davincible/gotiktoklive)

# GoTikTokLive

A Go module to download livestreams and discover, receive and decode livestreams and the livestream events such as comments and gifts in realtime from [TikTok LIVE](https://www.tiktok.com/live) by connecting to TikTok's internal WebCast push service. The package includes a wrapper that connects to the WebCast service using just the username (`uniqueId`). This allows you to connect to your own live chat as well as the live chat of other streamers. No credentials are required. Besides [Chat Comments](#ChatEvent), other events such as [Members Joining](#UserEvent), [Gifts](#GiftEvent), [Viewers](#ViewersEvent), [Follows](#UserEvent), [Shares](#UserEvent), [Questions](#QuestionEvent), [Likes](#LikeEvent) and [Battles](#BattlesEvent) can be tracked.

Looking for a Python implementation of this library? Check out [TikTok-Live-Connector](https://github.com/isaackogan/TikTok-Live-Connector) by [**@isaackogan**](https://github.com/isaackogan)

Looking for a Node.js implementation of this library? Check out [TikTok-Livestream-Chat-Connector](https://github.com/zerodytrash/TikTok-Livestream-Chat-Connector) by [**@zerodytrash**](https://github.com/zerodytrash)

**NOTE:** This is not an official API, and is no way affiliated to or sponsored by TikTok.

Go rewrite of [zerodytrash/TikTok-Livestream-Chat-Connector](https://github.com/zerodytrash/TikTok-Livestream-Chat-Connector)

#### Overview

- [Getting started](#getting-started)
- [Methods](#methods)
- [Events](#events)
- [Examples](#examples)
- [Contributing](#contributing)

## Getting started

1. Install the package using the Go package manager

```sh
go get github.com/Davincible/gotiktoklive
```

2. Create your first chat connection

```go
// Create TikTok Instance
tiktok := gotiktoklive.NewTikTok()

// Track a TikTok user by username
live, err := tiktok.TrackUser("promobot.robots")
if err != nil {
    panic(err)
}

// Start downloading stream
// Make sure you have the ffmpeg binary installed, and present in your path.
if err := live.DownloadStream(); err != nil {
    panic(err)
}

// Receive livestream events through the live.Events channel
for event := range live.Events {
    switch e := event.(type) {

    // You can specify what to do for specific events. All events are listed below.
    case gotiktoklive.UserEvent:
        fmt.Printf("%T : %s %s\n", e, e.Event, e.User.Username)

    // List viewer count
    case gotiktoklive.ViewersEvent:
        fmt.Printf("%T : %d\n", e, e.Viewers)

    // Specify the action for all remaining events
    default:
        fmt.Printf("%T : %+v\n", e, e)
    }
}
```

## Methods

```go
// TikTok allows you to track and discover current live streams.
type TikTok struct {

	Debug bool

	// LogRequests when set to true will log all made requests in JSON to debugHandler
	LogRequests bool
}

// NewTikTok creates a tiktok instance that allows you to track live streams
//  and discover current livestreams.
func NewTikTok() *TikTok {}

// TrackUser will start to track the livestream of a user, if live.
// To listen to events emitted by the livestream, such as comments and viewer
//  count, listen to the Live.Events channel.
// It will start a go routine and connect to the tiktok websocket.
func (t *TikTok) TrackUser(username string) (*Live, error) {}

// TrackRoom will start to track a room by room ID.
// It will start a go routine and connect to the tiktok websocket.
func (t *TikTok) TrackRoom(roomId string) (*Live, error) {}

// GetUserInfo will fetch information about the user, such as follwers stats,
//  their user ID, as well as the RoomID, with which you can tell if they are live.
func (t *TikTok) GetUserInfo(user string) (*UserInfo, error) {}

// GetRoomInfo will only fetch the room info, normally available with Live.Info
//  but not start tracking a live stream.
func (t *TikTok) GetRoomInfo(username string) (*RoomInfo, error) {}

// GetPriceList fetches the price list of tiktok coins. Prices will be given in
//  USD cents and the cents equivalent of the local currency of the IP location.
// To fetch a different currency, use a VPN or proxy to change your IP to a
//  different country.
func (t *TikTok) GetPriceList() (*PriceList, error) {}

// NewFeed creates a new Feed instance. Start fetching reccomended livestreams
//  with Feed.Next().
func (t *TikTok) NewFeed() *Feed {}

func (t *TikTok) SetDebugHandler(f func(...interface{})) {}

func (t *TikTok) SetErrorHandler(f func(error)) {}

func (t *TikTok) SetInfoHandler(f func(...interface{})) {}

func (t *TikTok) SetWarnHandler(f func(...interface{})) {}

// SetProxy will set a proxy for both the http client as well as the websocket.
// You can manually set a proxy with this method, or by using the HTTPS_PROXY
//  environment variable.
// ALL_PROXY can be used to set a proxy only for the websocket.
func (t *TikTok) SetProxy(url string, insecure bool) error {}
```

## Events

- [`RoomEvent`](#RoomEvent)
- [`ChatEvent`](#ChatEvent)
- [`UserEvent`](#UserEvent)
- [`ViewersEvent`](#ViewersEvent)
- [`GiftEvent`](#GiftEvent)
- [`LikeEvent`](#LikeEvent)
- [`QuestionEvent`](#QuestionEvent)
- [`ControlEvent`](#ControlEvent)
- [`MicBattleEvent`](#MicBattleEvent)
- [`BattlesEvent`](#BattlesEvent)
- [`RoomBannerEvent`](#RoomBannerEvent)
- [`IntroEvent`](#IntroEvent)

### RoomEvent

Room events are messages broadcast in the room. The most common event, is the
`Type: SystemMessage` at broadcast the beginning a stream gets watched, saying "Welcome to TikTok LIVE! Have fun interacting with others in real-time an
d remember to follow our Community Guidelines."

```go
type RoomEvent struct {
	Type    string
	Message string
}
```

### ChatEvent

Chat events are broadcasted when a user posts a chat message, aka comment to a livestream.

```go
type ChatEvent struct {
	Comment   string
	User      *User
	Timestamp int64
}
```

### UserEvent

User events are used when a user either joins the stream, shares the stream, or 
follows the host.

```go
type UserEvent struct {
	Event userEventType
	User  *User
}

type User struct {
	ID              int64
	Username        string
	Nickname        string
	ProfilePicture  *ProfilePicture
	ExtraAttributes *ExtraAttributes
	Badge           *BadgeAttributes
}

// User Event Types
const (
	USER_JOIN   userEventType = "user joined the stream"
	USER_SHARE  userEventType = "user shared the stream"
	USER_FOLLOW userEventType = "user followed the host"
)
```

### ViewersEvent

Viewer events broadcast the current amount of users watching the livestream.

```go
type ViewersEvent struct {
	Viewers int
}
```

### GiftEvent

Gift events are broadcast when a user buys a gift for the host.
To get more information about the gift, such as the price in coins,
find the gift by ID in the `live.GiftInfo.Gifts`.

Gift events with `GiftEvent.Type == 1` are streakable, meaning multiple gifts can 
be sent in sequence, such as roses. For these sequences, multiple events are broadcast.
Upon every subsequent gift in the streak, the `GiftEvent.RepeatCount` will increase by one.
To prevent the duplicate processing of streakable gifts, you should only process 
`if event.Type == 1 && event.RepeatEnd`, as this will be the final message of the streak, 
and includes the total number of gifts sent in the streak.

```go
type GiftEvent struct {
	ID          int
	Name        string
	Describe    string
	Cost        int
	RepeatCount int
	RepeatEnd   bool
	Type        int
	ToUserID    int64
	Timestamp   int64
	User        *User
}

type User struct {
	ID              int64
	Username        string
	Nickname        string
	ProfilePicture  *ProfilePicture
	ExtraAttributes *ExtraAttributes
	Badge           *BadgeAttributes
}
```

### LikeEvent

Like events are broadcast when a user likes the livestream. The event includes
the number of likes the user sent, as well as new total number of likes.

```go
type LikeEvent struct {
	Likes       int
	TotalLikes  int
	User        *User
	DisplayType string
	Label       string
}
```

### QuestionEvent

Question events are emitted when a question has been posted by a user. It includes
the question text, and the user that posted the question.

```go
type QuestionEvent struct {
	Quesion string
	User    *User
}
```

### ControlEvent

Control events are used to broadcast the status of the livestream.

Action values:

- 3: live stream ended

```go
type ControlEvent struct {
	Action int
}
```

### MicBattleEvent

```go
type MicBattleEvent struct {
	Users []*User
}
```

### BattlesEvent

```go
type BattlesEvent struct {
	Status  int
	Battles []*Battle
}

type Battle struct {
	Host   int64
	Groups []*BattleGroup
}

type BattleGroup struct {
	Points int
	Users  []*User
}
```

### RoomBannerEvent

Room banner event contains the JSON data unmarshaled into an interface that was
passed with the message.

```go
type RoomBannerEvent struct {
	Data interface{}
}
```

### IntroEvent

Intro events are broadcast upon connecting to a livestream.

```go
type IntroEvent struct {
	ID    int
	Title string
	User  *User
}
```

## Examples

### Fetching Recommended Live Streams

With a feed instance you can fetch a list of recommended livestreams, and directly
start tracking them with a single call.

```go
tiktok := NewTikTok()
feed := tiktok.NewFeed()

// Fetch 5 pages of recommended streams, usually 6 are returned at a time
for i := 0; i < 5; i++ {
	feedItem, err := feed.Next()
	if err != nil {
		panic(err)
	}

	for _, stream := range feedItem.LiveStreams {
		fmt.Printf("%s : %d viewers\n", stream.Room.Owner.Nickname, stream.Room.UserCount)
	}

	if !feedItem.Extra.HasMore {
		break
	}
}

recommendedStreams := feed.LiveStreams

// Start tracking the first stream
live, err := recommendedStreams[0].Track()
if err != nil {
	panic(err)
}

// Process events
...

```

### Error Handling

Gotiktoklive uses Go routines to fetch events using either websockets or HTTP polling.
These go routines need an error hander, that defaults to panic. You can overwrite this
behavior:

```go
tiktok.SetErrorHandler(func(err error) {
    ...
  })

```

## Contributing

Your improvements are welcome! Feel free to open an <a href="https://github.com/Davincible/gotiktoklive/issues">issue</a> or <a href="https://github.com/Davincible/gotiktoklive/pulls">pull request</a>.
