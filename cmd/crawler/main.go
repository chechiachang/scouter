package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	"github.com/chechiachang/scouter"
	"golang.org/x/oauth2"
)

func main() {
	githubApiToken := flag.String("token", "", "github api token (string)")
	flag.Parse()

	if *githubApiToken == "" {
		panic("Github api token is empty.")
	}

	log.Println("crawling...")

	// Prepare github client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *githubApiToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	layout := "2006-01-01"
	// set fetching with time range from start time to now
	endTime := time.Now()
	startTime, err := time.Parse(layout, "2008-01-01")
	if err != nil {
		panic(err)
	}

	// set fetch batch time interval
	startCursor := startTime
	endCursor := startCursor.AddDate(0, 1, 0) // interval: 1 month
	sort := "joined"
	order := "asc"

	for endCursor.Before(endTime) {
		wg := sync.WaitGroup{}
		defer wg.Done()

		query := "location:Taiwan created:" + startCursor.Format(layout) + ".." + endCursor.Format(layout)

		// First fetch
		r, err := scouter.SearchGithubUsers(tc, 1, query, sort, order)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Fetching ", query, ". Found records:", *r.Total)

		// paging fetch if result.Total > searchMaxPerPage
		if *r.Total > scouter.SearchMaxPerPage {

			pageNum := *r.Total / scouter.SearchMaxPerPage

			for page := 1; page < pageNum+1; page++ {
				pagedResult, err := scouter.SearchGithubUsers(tc, page, query, sort, order)
				time.Sleep(2 * time.Second)
				if err != nil {
					log.Fatal(err)
				}

				if err := scouter.UpsertGithubUsers(tc, pagedResult.Users); err != nil {
					log.Fatal(err)
				}

			}
		} else {
			if err := scouter.UpsertGithubUsers(tc, r.Users); err != nil {
				log.Fatal(err)
			}
		}

		// Move cursor forward 1 month
		startCursor = endCursor.AddDate(0, 0, 1)
		endCursor = startCursor.AddDate(0, 1, 0) // interval: 1 month

		wg.Wait()
	}
}
