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
		panic("Github api token is empty.")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *githubApiToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	log.Println("crawling...")

	query := "location:Taiwan"
	sort := "joined"
	order := "asc"
	r, err := scouter.SearchGithubUsers(tc, 1, query, "joined", "asc")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Found records:", *r.Total)

	pageNum := *r.Total / scouter.SearchMaxPerPage

	for page := 1; page < pageNum+1; page++ {
		result, err := scouter.SearchGithubUsers(tc, page, query, sort, order)
		if err != nil {
			log.Fatal(err)
		}

		if err := scouter.GetAvatar(result); err != nil {
			log.Fatal(err)
		}

		if err := scouter.UpsertUsers(result.Users); err != nil {
			log.Fatal(err)
		}

		// Github search API max rate is 30 queries/min for authorized user
		// No need. The avatar downloading require more than 2 sec
		// time.Sleep(2 * time.Second)
	}
}
