package scouter

import (
	"context"
	"github.com/google/go-github/github"
	"log"
)

func GetUser(username string) (*github.User, error) {
	client := github.NewClient(nil)

	user, resp, err := client.Users.Get(context.Background(), username)
	log.Print(resp)

	return user, err
}

func FetchUsers(username string) ([]*github.User, error) {
	client := github.NewClient(nil)

	opt := &github.UserListOptions{}

	users, resp, err := client.Users.ListAll(context.Background(), opt)
	log.Println(resp)

	return users, err
}
