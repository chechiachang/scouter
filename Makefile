.PHONY: crawler
crawler:
	go build ./cmd/crawler

avatar_downloader:
	go build ./cmd/avatar_downloader

.PHONY: test
test:
	go test ./...

.PHONY: apiserver
apiserver:
	go build ./cmd/apiserver
