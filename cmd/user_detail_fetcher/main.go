package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"runtime"
	"time"

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

	if err := updateUsersDetail(tc); err != nil {
		log.Fatal(err)
	}

}
func updateUsersDetail(tc *http.Client) error {
	log.Println("Starting upsert db user with github user api...")

	total, err := scouter.CountUsers()
	if err != nil {
		return err
	}

	pageSize := scouter.SearchMaxPerPage
	pageNum := total / pageSize
	sort := "$natural"

	runtime.GOMAXPROCS(1)

	for page := 1; page < pageNum+1; page++ {

		log.Println("Paging ", page, "/", pageNum)
		users, err := scouter.FindUsers(bson.M{}, sort, page, pageSize)
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
