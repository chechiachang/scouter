package main

import (
	"fmt"
	"github.com/chechiachang/scouter"
)

func main() {
	fmt.Println("crawling...")

	users, err := scouter.FetchUsers("chechiachang")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
