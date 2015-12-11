NAME := cacheout

test: unit acceptance lint

acceptance: build
	bats test

build: dependencies
	go build -o bin/$(NAME)

unit: dependencies
	go test ./...

lint: dependencies
	@which -s gometalinter || (go get github.com/alecthomas/gometalinter && gometalinter --install)
	gometalinter

dependencies:
	go get -t

.PHONY: acceptance build dependencies lint test unit
