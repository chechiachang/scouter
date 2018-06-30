package scouter

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

const (
	UserMaxPerPage   = 100
	SearchMaxPerPage = 1000
)

func GetUser(username string) (*github.User, error) {
	client := github.NewClient(nil)

	user, resp, err := client.Users.Get(context.Background(), username)
	log.Print(resp)

	return user, err
}

func FetchUsers() ([]*github.User, error) {
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
func SearchUsers() (*github.UsersSearchResult, error) {
	client := github.NewClient(nil)

	opt := &github.SearchOptions{
		Sort:      "joined",
		Order:     "asc",
		TextMatch: false,
		ListOptions: github.ListOptions{
			Page:    0,
			PerPage: SearchMaxPerPage,
		},
	}

	query := "location:Taiwan"
	result, resp, err := client.Search.Users(context.Background(), query, opt)
	log.Println(resp)

	return result, err
}
