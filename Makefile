CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-fetch
	cp *.go src/github.com/whosonfirst/go-whosonfirst-fetch/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-cli"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-readwrite-bundle"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-geojson-v2"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-index"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-csv"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-log"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"
	mv src/github.com/whosonfirst/go-whosonfirst-index/vendor/github.com/whosonfirst/go-whosonfirst-sqlite src/github.com/whosonfirst/
	mv src/github.com/whosonfirst/go-whosonfirst-readwrite-bundle/vendor/github.com/whosonfirst/go-whosonfirst-readwrite src/github.com/whosonfirst/
	rm -rf src/github.com/whosonfirst/go-whosonfirst-index/vendor/github.com/whosonfirst/go-whosonfirst-log
	rm -rf src/github.com/whosonfirst/go-whosonfirst-readwrite-bundle/vendor/github.com/whosonfirst/go-whosonfirst-cli
	rm -rf src/github.com/whosonfirst/go-whosonfirst-readwrite-bundle/vendor/github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/vendor/github.com/whosonfirst/go-whosonfirst-sqlite
	rm -rf src/github.com/whosonfirst/go-whosonfirst-readwrite-bundle/vendor/github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/vendor/github.com/whosonfirst/go-whosonfirst-sqlite-features/vendor/github.com/whosonfirst/go-whosonfirst-index

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt *.go

bin: 	self
	rm -rf bin/*
	GOPATH=$(GOPATH) go build -o bin/wof-fetch cmd/wof-fetch.go

# leaving these here as a reference but they will _not_ work (as in compile correctly)
# because Go won't know where to find the requisite sqlite hoohah for other platforms
# (20181030/thisisaaronland)

dist-build:
	OS=darwin make dist-os
	OS=windows make dist-os
	OS=linux make dist-os

dist-os:
	if test ! -d dist/$(OS); then mkdir -p dist/$(OS); fi
	GOOS=$(OS) GOPATH=$(GOPATH) GOARCH=386 go build -o dist/$(OS)/wof-fetch cmd/wof-fetch.go
	chmod +x dist/$(OS)/wof-fetch
	cd dist/$(OS) && shasum -a 256 wof-fetch > wof-fetch.sha256
