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

.PHONY: build
build: user_fetcher user_detail_fetcher contribution_fetcher avatar_downloader apiserver

.PHONY: test
test:
	go test ./...

.PHONY: test
encoding:
	rm -f data/encodings data/index
	python ./face_recognition/encoding_file_generator.py

.PHONY: apiserver
apiserver:
	python ./face_recognition/apiserver.py
