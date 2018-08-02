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

go build ./cmd/crawler/ && ./crawler --token github-api-token

# Thanks

### Open source packages
go-github

### Issues

```
GET https://api.github.com/search/users?order=asc&page=11&per_page=1000&q=location%3ATaiwan&sort=joined: 422 Only the first 1000 search results are available []
```
