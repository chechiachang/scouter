scouter
---


[![Build Status](https://travis-ci.org/chechiachang/scouter.svg?branch=master)](https://travis-ci.org/chechiachang/scouter)


# TODOs

- [] Github api crawler
  - [v] Add an api to fetch user list
    - [v] order by most follower
    - [v] order by most commit. Might need query by username.
  - [v] Save user data to mongodb
    - [v] username
    - [v] avatar
    - [] # of follower
    - [] # of commits
- [] Face Recognizer
  - [] Face detector
    - [] Generate face landmarks
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

