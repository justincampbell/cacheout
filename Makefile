NAME := cacheout
HOMEPAGE := https://github.com/justincampbell/$(NAME)
PREFIX := /usr/local

VERSION=0.1.0
TAG=v$(VERSION)

ARCHIVE=$(NAME)-$(TAG).tar.gz
ARCHIVE_URL=$(HOMEPAGE)/archive/$(TAG).tar.gz

test: unit acceptance lint

release: tag sha

tag:
	git tag --force latest
	git tag | grep $(TAG) || git tag --message "Release $(TAG)" --sign $(TAG)
	git push origin
	git push origin --force --tags

pkg/$(ARCHIVE): pkg/
	wget --output-document pkg/$(ARCHIVE) $(ARCHIVE_URL)

pkg/:
	mkdir pkg

sha: pkg/$(ARCHIVE)
	shasum pkg/$(ARCHIVE)

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/$(NAME) $(PREFIX)/bin/$(NAME)

uninstall:
	rm -vf $(PREFIX)/bin/$(NAME)

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

.PHONY: acceptance build dependencies install release sha tag test uninstall unit
