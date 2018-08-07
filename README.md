scouter
---

[![Build Status](https://travis-ci.org/chechiachang/scouter.svg?branch=master)](https://travis-ci.org/chechiachang/scouter)

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
- [] Google Search API Face downloader
  - [] Search Avatar with github username and login
  - [] [Google Custom Search API](https://developers.google.com/custom-search/docs/tutorial/introduction)
- [v] Face Recognizer
  - [v] Face detector
    - [] Generate face landmarks
    - [v] [Face recognition](https://github.com/ageitgey/face_recognition)
- [] Front-End
  - [v] Unity ios app
  - [] API portal
  - [] AR
- [] Readme
  - [] Add a pip requirements

# How to use

```
go build ./cmd/crawler/ && ./crawler --token github-api-token
```

### Face detection

[Face Detection](https://github.com/ageitgey/face_recognition)

```
face_recognition --show-distance true --tolerance 0.54 ./pictures_of_people_i_know/ ./unknown_pictures/
```

apiserver
```
pip3 install dlib flask face_recognition pymongo bson
```

# Thanks

### Open source packages
go-github

### Issues

```
GET https://api.github.com/search/users?order=asc&page=11&per_page=1000&q=location%3ATaiwan&sort=joined: 422 Only the first 1000 search results are available []
```

2604

# Develop on Mac

Scouter is developed on Mac. As a new developer on ios, I was blocked by lots of envronment setting on unity and xcode.
Here are things to setup a develop environments.

# Golang

user data fetcher

# Unity

- C# 
- .Net
- [dlibDotNet](https://github.com/takuya-takeuchi/DlibDotNet)

https://github.com/Unity-Technologies/EntityComponentSystemSamples/issues/31

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
