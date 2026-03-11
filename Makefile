NAME := cacheout

test: unit acceptance

acceptance: build
	bats test

build:
	go build -o bin/$(NAME)

unit:
	go test ./...

lint:
	@which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

.PHONY: acceptance build lint test unit
