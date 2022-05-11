build:
	go build -o qa-server

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o qa-server

.PHONY: build build-linux
