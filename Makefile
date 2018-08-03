.PHONY: crawler
crawler:
	go build ./cmd/crawler

.PHONY: avatar_downloader
avatar_downloader:
	go build ./cmd/avatar_downloader

.PHONY: apiserver
apiserver:
	go build ./cmd/apiserver

.PHONE: build
build: crawler avatar_downloader apiserver

.PHONY: test
test:
	go test ./...

