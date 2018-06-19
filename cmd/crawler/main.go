package main

import (
	"github.com/chechiachang/scouter"
	"log"
)

func main() {
	log.Println("crawling...")

	users, err := scouter.FetchUsers("chechiachang")
	if err != nil {
		log.Println(err)
	}
	log.Println(users)
}
