package main

import (
	"log"

	"github.com/chechiachang/scouter"
)

func main() {
	log.Println("crawling...")

	//users, err := scouter.FetchUsers()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(users)

	result, err := scouter.SearchUsers()
	if err != nil {
		log.Println(err)
	}
	log.Println(*result.Total)

	err = scouter.InsertUsers(result.Users)
	if err != nil {
		log.Fatal(err)
	}
}
