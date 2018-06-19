.PHONY: crawler
crawler:
	go build ./cmd/crawler && ./crawler

test:
	go test ./...

apiserver:
	go build ./cmd/apiserver
