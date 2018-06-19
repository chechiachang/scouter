package main

import (
	"fmt"
)

func main() {
	fmt.Println("crawling...")

	users, err := FetchUsers("chechiachang")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
