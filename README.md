Scouter
===

Scouter: A human face detector which show your github contribution statistics.

[![Build Status](https://travis-ci.org/chechiachang/scouter.svg?branch=master)](https://travis-ci.org/chechiachang/scouter)

# Fetching data from github with github api

1. Have a local running mongoDB

```
docker run -d --name mongo mongo
```

2. Generate github access token
  - User -> settings -> Developer settings -> Personal access tokens
  - Keep your token safe

3. Run fetchers with token

```
# Fetch user with Github Search API
go build ./cmd/user_fetcher && ./user_fetcher -token <github-api-token>

# Fetch user detail information like follwers and repos with Github User API
go build ./cmd/user_detail_fetcher && ./user_detail_fetcher -token <github-api-token>

# Fetch users' avatar with user.url from data in mongodb
go build ./cmd/avatar_downloader && ./avatar_downloader

# Fetch users' contribution statics by parsing html response of user.url from data in mongodb
go build ./cmd/contribution_fetcher && ./contribution_fetcher
```

4. Make sure data are good to go

Check user data in mongodb
```
docker exec -it mongo scouter
db.users.findOne()
```

Check users' avatar
```
ls data/avatars
```

# Face detection and Face recognition

[Face Recognition API](https://github.com/ageitgey/face_recognition)

1. Install python dependency
```
pip3 install dlib flask face_recognition pymongo bson
```

1. Try some face recognition api
```
face_recognition --show-distance true --tolerance 0.54 ./pictures_of_people_i_know/ ./unknown_pictures/
```

2. Prepare face identity file with encoding generator
```
# Filter data/avatars image. Save images with human faces to data/human_face.
# Generate face_recognition/encodings and face_recognition/index with data/human_face
python ./face_recognition/encoding_file_generator.py
```

3. Run our apiserver
```
python ./face_recognition/apiserver.py
```

# Unity

### Warning

1. Some of the contents are priced.
2. Some of the code in this part are heavily broken lol.

### Use unity

1. Have a working unity and unity account

2. Create a new project

3. Download and import a free face tracker example from unity asset store
[Face Tracker Example](https://assetstore.unity.com/packages/templates/tutorials/facetracker-example-35284)

4. Download and import another priced asset is required: 
[OpenCV for Unity](https://assetstore.unity.com/packages/tools/integration/opencv-for-unity-21088)
NOTE: This is a priced asset.

5. Copy unity scenes and scripts
```
unity/Assets/Scouter/* to /Users/Shared/Unity/<your-project>/Assets/FaceTrackerExample
```

6. Test run in unity

### Project Configuration

1. Open unity build for iphone project with xcode. Open another project.
2. Xcode developer account:
  - Xcode - Preferences - Accounts: Add and login your 'apple developer ID'. 
  - The team of your developer account will show up. 
  - In my case, A personal team show up with my username as team name.
3. Signing:
  - Click my-project. The project configure page will show up.
  - General - Identity: Change your display name and Bundle Identifier. Any reasonable identifier other than the example identifier will work.
  - General - Signing: Check 'Automatically manage signing'.
  - Choose your team. A signing certificates will show up.
  - If you stuck here, check your bundle identifier.

### Build project

1. Attach your device (your iphone). Unlock your iphone.
2. Click 'Build and Run Current Schema'.

# TODOs

- [v] Github api crawler
  - [v] Add an api to search user in Taiwan 
    - [v] Order by joined asc
    - [v] Implement a api call with narrowed search condition
  - [v] Add an api to fetch user Data
    - [v] Get user with userUrl
    - [v] order by most follower
    - [v] order by most commit. Might need query by username.
  - [v] Save user data to mongodb
    - [v] username
    - [v] avatar
    - [v] # of follower
    - [v] # of contributions
- [x] Google Search API Face downloader
  - [x] Search Avatar with github username and login
  - [x] Google Custom Search API
    [Google Custom Search API](https://developers.google.com/custom-search/docs/tutorial/introduction)
- [v] Face Recognizer
  - [v] Face detector
    - [v] Generate face encoding and save to Python Pickle file
    - [v] Face Recognition
      [Face recognition](https://github.com/ageitgey/face_recognition)
- [v] Front-End
  - [v] Unity ios app
  - [v] API portal
  - [ ] AR GUI
- [v] Readme


