.PHONY: crawler
crawler:
	go build ./cmd/crawler

.PHONY: test
test:
	go test ./...

.PHONY: apiserver
apiserver:
	go build ./cmd/apiserver
