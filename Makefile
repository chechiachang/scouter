.PHONY: crawler
crawler:
	go build ./cmd/crawler && ./crawler

.PHONY: test
test:
	go test ./...

.PHONY: apiserver
apiserver:
	go build ./cmd/apiserver
