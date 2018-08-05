package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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
		users, err := scouter.FindUsers(bson.M{}, sort, page, pageSize)
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
					str := contributionLine.FindString(content) // 1,353 contributions or 1 contribution.
					str = strings.Replace(str, ",", "", -1)     // 1,353 or 1. Remove comma.
					str = contributionNumber.FindString(str)    // 1353 or 1.

					if str != "" {
						c, err := strconv.Atoi(str) // 0
						if err != nil {
							fmt.Println(err)
						}
						contribution += c
					}
				})
			}

			user.Contribution = contribution
			fmt.Printf("User: %d %s %d \n", user.ID, user.GetLogin(), user.Contribution)

			// update user
			if err := scouter.UpsertUser(user); err != nil {
				return err
			}

		}

	}

	return nil
}
