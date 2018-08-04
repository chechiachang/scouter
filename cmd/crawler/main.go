package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chechiachang/scouter"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/oauth2"
)

func main() {
	githubApiToken := flag.String("token", "", "github api token (string)")
	flag.Parse()

	if *githubApiToken == "" {
		panic("Github api token is empty.")
	}

	log.Println("Starting crawler...")

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

	if err := countContribution(); err != nil {
		log.Fatal(err)
	}

}

func searchUsers(tc *http.Client) error {
	log.Println("Starting fetch github user with search api...")

	layout := "2006-01-01T00:00:00"
	// set fetching with time range from start time to now
	endTime := time.Now()
	startTime, err := time.Parse(layout, "2008-01-01T00:00:00")
	if err != nil {
		return err
	}

	total, err := scouter.CountGithubUsers(tc, "location:taiwan")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Total: ", total)

	// set fetch batch time interval
	startCursor := startTime
	endCursor := startCursor.AddDate(0, 1, 0) // interval: 1 month
	sort := "joined"
	order := "asc"

	runtime.GOMAXPROCS(1)

	for endCursor.Before(endTime) {
		query := "location:Taiwan created:" + startCursor.Format(layout) + ".." + endCursor.Format(layout)

		// First fetch
		r, err := scouter.SearchGithubUsers(tc, 1, query, sort, order)
		time.Sleep(2 * time.Second) // Github search API max rate
		if err != nil {
			return err
		}
		log.Println("Fetching ", query, ". Found records:", r.GetTotal())

		// paging fetch if result.Total > searchMaxPerPage
		if *r.Total > scouter.SearchMaxPerPage {

			log.Fatal("Pagesize exceed ", scouter.SearchMaxPerPage, ". Some data may not be fetched")

		} else {

			for _, user := range r.Users {
				u := scouter.User{
					ID:   user.GetID(),
					User: &user,
				}
				if err := scouter.UpsertUser(u); err != nil {
					return err
				}
			}
		}

		// Move cursor forward 1 month
		startCursor = startCursor.AddDate(0, 1, 0)
		endCursor = endCursor.AddDate(0, 1, 0) // interval: 1 month

	}
	return nil
}

func updateUsersDetail(tc *http.Client) error {
	log.Println("Starting upsert db user with github user api...")

	total, err := scouter.CountUsers()
	if err != nil {
		return err
	}

	pageSize := scouter.SearchMaxPerPage
	pageNum := total / pageSize

	runtime.GOMAXPROCS(1)

	for page := 1; page < pageNum+1; page++ {

		log.Println("Paging ", page, "/", pageNum)
		users, err := scouter.FindUsers(bson.M{}, page, pageSize)
		if err != nil {
			return err
		}

		for _, user := range users {

			detailedUser, err := scouter.GetGithubUser(tc, user.GetLogin())
			time.Sleep(750 * time.Microsecond) // Github search API max rate per query
			if err != nil {
				return err
			}

			u := scouter.User{
				ID:   detailedUser.GetID(),
				User: detailedUser,
			}

			if err := scouter.UpsertUser(u); err != nil {
				return err
			}
		}
	}
	return nil
}

// Parse this
// https://github.com/users/winson/contributions?from=2017-12-01&to=2017-12-31&full_graph=1
func countContribution() error {
	log.Println("Starting counting contribution with github user ...")

	githubUserUrl := "https://github.com/users/%s/contributions?from=%d-01-01&to=%d-12-31&full_graph=1"

	total, err := scouter.CountUsers()
	if err != nil {
		return err
	}

	pageSize := scouter.SearchMaxPerPage
	pageNum := total / pageSize

	r, err := regexp.Compile(".[0-9]* contributions")
	if err != nil {
		return err
	}

	for page := 1; page < pageNum+1; page++ {

		log.Println("Paging ", page, "/", pageNum)
		users, err := scouter.FindUsers(bson.M{}, page, pageSize)
		if err != nil {
			return err
		}

		for _, user := range users {

			contribution := 0

			// Count contribution 2008..2018
			for year := 2008; year < 2019; year++ {
				url := fmt.Sprintf(githubUserUrl, user.GetLogin(), year, year)
				res, err := http.Get(url)
				if err != nil {
					return err
				}
				defer res.Body.Close()

				//get contribution count from graph
				doc, err := goquery.NewDocumentFromReader(res.Body)
				if err != nil {
					return err
				}

				// Get contiribution number from h2 content
				doc.Find(".js-contribution-graph .text-normal").Each(func(i int, s *goquery.Selection) {
					content := s.Text()
					result := r.FindString(content)                                        // 0 contributions
					c, _ := strconv.Atoi(strings.Replace(result, " contributions", "", 1)) // 0
					contribution += c
				})
			}

			user.Contribution = contribution
			fmt.Println(user.Contribution)

			// update user
			if err := scouter.UpsertUser(user); err != nil {
				return err
			}

		}

	}

	return nil
}
