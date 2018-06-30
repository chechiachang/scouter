package main

import (
	"context"
	"flag"
	"log"

	"github.com/chechiachang/scouter"
	"golang.org/x/oauth2"
)

func main() {
	githubApiToken := flag.String("token", "", "github api token (string)")
	flag.Parse()

	if *githubApiToken == "" {
		panic("Token is empty.")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *githubApiToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	log.Println("crawling...")

	//users, err := scouter.FetchUsers()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(users)

	result, err := scouter.SearchUsers(tc)
	if err != nil {
		log.Println(err)
	}
	log.Println(*result.Total)

	err = scouter.InsertUsers(result.Users)
	if err != nil {
		log.Fatal(err)
	}
}
