GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: proto
proto:
	@protoc \
		--proto_path=. \
		--go_out=:. \
		proto/tiktok.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./... -cover -coverprofile coverage.out -count=1

.PHONY: cov
cov:
	@gocovsh --profile coverage.out


