scouter
---

[![Build Status](https://travis-ci.org/chechiachang/scouter.svg?branch=master)](https://travis-ci.org/chechiachang/scouter)

# TODOs

- [] Github api crawler
  - [v] Add an api to search user in Taiwan 
    - [v] Order by joined asc
    - [ ] Implement a api call with narrowed search condition
  - [v] Add an api to fetch user Data
    - [] Get user with userUrl
    - [] order by most follower
    - [] order by most commit. Might need query by username.
  - [v] Save user data to mongodb
    - [v] username
    - [v] avatar
    - [] # of follower
    - [] # of commits
- [] Google Search API Face downloader
  - [] Search Avatar with github username and login
  - [] [Google Custom Search API](https://developers.google.com/custom-search/docs/tutorial/introduction)
- [] Face Recognizer
  - [] Face detector
    - [] Generate face landmarks
    - [] [Face recognition](https://github.com/ageitgey/face_recognition)
  - [] Face modle trainer
  - [] Face recognizer
- [] Front-End
  - [] API portal
  - [] AR

# How to use

```
pip3 install dlib
pip3 install face_recognition
```

```
go build ./cmd/crawler/ && ./crawler --token github-api-token
```

# Thanks

### Open source packages
go-github

### Issues

```
GET https://api.github.com/search/users?order=asc&page=11&per_page=1000&q=location%3ATaiwan&sort=joined: 422 Only the first 1000 search results are available []
```

# Develop on Mac

Scouter is developed on Mac. As a new developer on ios, I was blocked by lots of envronment setting on unity and xcode.
Here are things to setup a develop environments.

# Golang

# Unity

# Xcode

Version: 9.4.1

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
