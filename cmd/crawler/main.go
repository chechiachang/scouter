package main

import (
	"context"
	"flag"
	"log"
	"net/http"
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

	if err := searchUsers(tc); err != nil {
		log.Fatal(err)
	}

	if err := updateUsersDetail(tc); err != nil {
		log.Fatal(err)
	}
}

func searchUsers(tc *http.Client) error {
	layout := "2006-01-01"
	// set fetching with time range from start time to now
	endTime := time.Now()
	startTime, err := time.Parse(layout, "2008-01-01")
	if err != nil {
		return err
	}

	// set fetch batch time interval
	startCursor := startTime
	endCursor := startCursor.AddDate(0, 1, 0) // interval: 1 month
	sort := "joined"
	order := "asc"

	wg := sync.WaitGroup{}
	defer wg.Done()

	for endCursor.Before(endTime) {

		query := "location:Taiwan created:" + startCursor.Format(layout) + ".." + endCursor.Format(layout)

		// First fetch
		r, err := scouter.SearchGithubUsers(tc, 1, query, sort, order)
		time.Sleep(2 * time.Second) // Github search API max rate
		if err != nil {
			return err
		}
		log.Println("Fetching ", query, ". Found records:", *r.Total)

		// paging fetch if result.Total > searchMaxPerPage
		if *r.Total > scouter.SearchMaxPerPage {

			pageNum := *r.Total / scouter.SearchMaxPerPage

			for page := 1; page < pageNum+1; page++ {

				pagedResult, err := scouter.SearchGithubUsers(tc, page, query, sort, order)
				time.Sleep(2 * time.Second) // Github search API max rate
				if err != nil {
					return err
				}

				if err := scouter.UpsertUsers(pagedResult.Users); err != nil {
					return err
				}

			}

		} else {
			if err := scouter.UpsertUsers(r.Users); err != nil {
				return err
			}
		}

		// Move cursor forward 1 month
		startCursor = endCursor.AddDate(0, 0, 1)
		endCursor = startCursor.AddDate(0, 1, 0) // interval: 1 month

		//	time.Sleep(time.Duration(*r.Total*750) * time.Microsecond) // Github search API max rate per query
	}
	wg.Wait()
	return nil
}

func updateUsersDetail(tc *http.Client) error {
	return nil
}
