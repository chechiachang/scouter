package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	//"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/chechiachang/scouter"
	"github.com/globalsign/mgo/bson"
)

func main() {

	log.Println("Starting crawler...")

	if err := countContribution(); err != nil {
		log.Fatal(err)
	}

}

const (
	githubUserUrl = "https://github.com/users/%s/contributions?from=%d-01-01&to=%d-12-31&full_graph=1"
)

// Parse this
// https://github.com/users/winson/contributions?from=2017-12-01&to=2017-12-31&full_graph=1
func countContribution() error {
	log.Println("Starting counting contribution with github user ...")

	total, err := scouter.CountUsers()
	if err != nil {
		return err
	}

	pageSize := scouter.SearchMaxPerPage
	pageNum := total / pageSize

	// Use regex to get contribution number
	contributionLine, err := regexp.Compile("[0-9,]* contribution")
	if err != nil {
		return err
	}
	contributionNumber, err := regexp.Compile("[0-9,]*")
	if err != nil {
		return err
	}

	// paging data from mongo db
	for page := 1; page < pageNum+1; page++ {

		log.Println("Paging ", page, "/", pageNum)
		sort := "$natural"
		users, err := scouter.FindUsers(bson.M{"user.type": "User"}, sort, page, pageSize)
		if err != nil {
			return err
		}

		for _, user := range users {
			sumContribution(user, contributionLine, contributionNumber)
		}
	}

	return nil
}

func sumContribution(user scouter.User, contributionLine, contributionNumber *regexp.Regexp) {

	wg := sync.WaitGroup{}
	contribution := make(chan int, 11)

	// Count contribution 2008..2019
	for year := 2008; year < 2020; year++ {

		go func() {
			wg.Add(1)
			defer wg.Done()

			url := fmt.Sprintf(githubUserUrl, user.GetLogin(), year, year)
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			//get contribution count from graph
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Get contiribution number from h2 content
			doc.Find(".js-contribution-graph .text-normal").Each(func(i int, s *goquery.Selection) {
				content := s.Text()
				str := contributionLine.FindString(content) // 1,353 contributions or 1 contribution.
				str = strings.Replace(str, ",", "", -1)     // 1,353 or 1. Remove comma.
				str = contributionNumber.FindString(str)    // 1353 or 1.

				if str != "" {
					c, err := strconv.Atoi(str) // 0
					if err != nil {
						log.Fatal(err)
					}
					contribution <- c
				}
			})
		}()

	}
	wg.Wait()

	// update user
	user.Contribution += <-contribution
	fmt.Printf("User: %d %s %d \n", user.ID, user.GetLogin(), user.Contribution)
	if err := scouter.UpsertUser(user); err != nil {
		log.Fatal(err)
	}
}
