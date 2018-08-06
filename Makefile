DOCKERHUB_USER="chechiachang"

# Build

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

# Test & Run

.PHONY: test
test:
	go test ./...

.PHONY: apiserver
apiserver:
	python ./face_recognition/apiserver.py

# Build & ship

.PHONY: encodings
encoding:
	rm -f face_recognition/encodings face_recognition/index
	python ./face_recognition/encoding_file_generator.py

.PHONY: base
base:
	time docker build \
    --tag $(DOCKERHUB_USER)/scouter-apiserver-base \
		--file face_recognition/base/Dockerfile .

.PHONY: image
image:
	time docker build \
    --tag $(DOCKERHUB_USER)/scouter-apiserver \
		--file face_recognition/Dockerfile .

.PHONY: unity
unity:
	cp /Users/Shared/Unity/scouter2/Assets/FaceTrackerExample/Scenes/WebCamTextureFaceTrackerExample.unity unity/Assets/Scouter/Scenes
	cp /Users/Shared/Unity/scouter2/Assets/FaceTrackerExample/Scenes/WebCamTextureFaceTrackerExample.unity.meta unity/Assets/Scouter/Scenes
	cp /Users/Shared/Unity/scouter2/Assets/FaceTrackerExample/Scripts/WebCamTextureFaceTrackerExample.cs unity/Assets/Scouter/Scripts
	cp /Users/Shared/Unity/scouter2/Assets/FaceTrackerExample/Scripts/WebCamTextureFaceTrackerExample.cs.meta unity/Assets/Scouter/Scripts

