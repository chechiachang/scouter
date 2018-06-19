.PHONY: crawler
crawler:
	go build ./src/cmd/crawler && ./crawler

apiserver:
	go build ./src/cmd/apiserver
