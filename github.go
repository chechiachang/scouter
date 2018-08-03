package scouter

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

const (
	UserMaxPerPage   = 100
	SearchMaxPerPage = 1000
)

func GetGithubUser(username string) (*github.User, error) {
	client := github.NewClient(nil)

	user, resp, err := client.Users.Get(context.Background(), username)
	log.Print(resp)

	return user, err
}

func FetchGithubUsers() ([]*github.User, error) {
	client := github.NewClient(nil)

	opt := &github.UserListOptions{

		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: UserMaxPerPage,
		},
	}

	users, resp, err := client.Users.ListAll(context.Background(), opt)
	log.Println(resp)

	return users, err
}

// https://developer.github.com/v3/search/#search-users
// location:Taiwan&sort=joined&order=asc
func SearchGithubUsers(tc *http.Client, page int, query, sort, order string) (*github.UsersSearchResult, error) {
	client := github.NewClient(tc)

	opt := &github.SearchOptions{
		Sort:      sort,
		Order:     order,
		TextMatch: false,
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: SearchMaxPerPage,
		},
	}

	result, resp, err := client.Search.Users(context.Background(), query, opt)
	log.Println(resp)

	return result, err
}