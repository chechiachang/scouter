.PHONY: user_fetcher
user_fetcher:
	go build ./cmd/user_fetcher
	
.PHONY: user_detail_fetcher
user_detail_fetcher:
	go build ./cmd/user_detail_fetcher

.PHONY: contribution_fetcher
contribution_fetcher:
	go build ./cmd/contribution_fetcher

.PHONY: avatar_downloader
avatar_downloader:
	go build ./cmd/avatar_downloader

.PHONY: apiserver
apiserver:
	go build ./cmd/apiserver

.PHONY: build
build: user_fetcher user_detail_fetcher contribution_fetcher avatar_downloader apiserver

.PHONY: test
test:
	go test ./...

