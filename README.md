Scouter
===

Scouter: A human face detector which show your Github contribution statistics.

[![Build Status](https://travis-ci.org/chechiachang/scouter.svg?branch=master)](https://travis-ci.org/chechiachang/scouter)

![inline](docs/2000.jpeg)

![inline](docs/demo.png)

# Brief

1. Fetch following data from Github with Github API
  - User data
  - User avatar
  - Parse HTML to get user contribution statistics
2. Convert avatar to identity with Face Recognition API. Encoding avatar identity with userId.
3. Track face and crop face image from camera streaming with OpenCV
4. Send face image to Flask API server
  - Convert unknown face image to identity.
  - Get userId and contribution statistics with identity.
5. Send user contribution to App and display.

# Fetching data from Github with Github API

1. Have a local running mongoDB

```
docker run -d --name mongo mongo
```

2. Generate Github access token
  - User -> settings -> Developer settings -> Personal access tokens
  - Keep your token safe

3. Run fetchers with token

```
# Fetch user with Github Search API
go build ./cmd/user_fetcher && ./user_fetcher -token <Github-API-token>

# Fetch user detail information like follwers and repos with Github User API
go build ./cmd/user_detail_fetcher && ./user_detail_fetcher -token <Github-API-token>

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

[Face Recognition API](https://Github.com/ageitgey/face_recognition)

1. Install python dependency
```
pip3 install dlib flask face_recognition pymongo bson
```

1. Try some face recognition API
```
face_recognition --show-distance true --tolerance 0.54 ./pictures_of_people_i_know/ ./unknown_pictures/
```

2. Prepare face identity file with encoding generator
```
# Filter data/avatars image. Save images with human faces to data/human_face.
# Generate face_recognition/encodings and face_recognition/index with data/human_face
python ./face_recognition/encoding_file_generator.py
```

3. Run our APIserver
```
python ./face_recognition/APIserver.py
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

- [x] Github API crawler
  - [x] Add an API to search user in Taiwan 
    - [x] Order by joined asc
    - [x] Implement a API call with narrowed search condition
  - [x] Add an API to fetch user Data
    - [x] Get user with userUrl
    - [x] order by most follower
    - [x] order by most commit. Might need query by username.
  - [x] Save user data to mongodb
    - [x] username
    - [x] avatar
    - [x] # of follower
    - [x] # of contributions
- [x] Google Search API Face downloader
  - [x] Search Avatar with Github username and login
  - [x] Google Custom Search API
    [Google Custom Search API](https://developers.google.com/custom-search/docs/tutorial/introduction)
- [x] Face Recognizer
  - [x] Face detector
    - [x] Generate face encoding and save to Python Pickle file
    - [x] Face Recognition
      [Face recognition](https://Github.com/ageitgey/face_recognition)
- [x] Front-End
  - [x] Unity ios app
  - [x] API portal
  - [ ] AR GUI
- [x] Readme


