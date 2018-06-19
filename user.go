package scouter

import (
	"context"
	"github.com/google/go-github/github"
)

func FetchUsers(username string) ([]*github.User, error) {
	client := github.NewClient(nil)

	opt := &github.UserListOptions{}

	users, _, err := client.Users.ListAll(context.Background(), opt)
	return users, err
}
