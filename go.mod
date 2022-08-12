module github.com/Davincible/gotiktoklive

go 1.17

require github.com/gobwas/ws v1.1.0

require (
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
)

require (
	github.com/chromedp/cdproto v0.0.0-20220812200530-d0d83820bffc
	github.com/chromedp/chromedp v0.8.3
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20220812174116-3211cb980234
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	google.golang.org/protobuf v1.28.1
)

replace github.com/Davincible/gotiktoklive => ./
