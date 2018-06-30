package scouter

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/go-github/github"
)

const (
	UserMaxPerPage   = 100
	SearchMaxPerPage = 1000
)

type User struct {
	ID int64 `bson:"_id" json:"id"`
	*github.User
}

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
func SearchUsers(tc *http.Client) (*github.UsersSearchResult, error) {
	client := github.NewClient(tc)

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

func GetAvatar(result *github.UsersSearchResult) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Println("work dir: " + wd)
	workspace := filepath.Join(wd, "data/avatars")
	if _, err := os.Stat(workspace); os.IsNotExist(err) {
		log.Println(workspace + " not found. Mkdir one.")
		if err = os.MkdirAll(workspace, 0644); err != nil {
			return err
		}
	}

	for _, user := range result.Users {
		dir := filepath.Join(workspace, strconv.FormatInt(*user.ID, 10))
		log.Println(dir)
		if err = os.MkdirAll(dir, 0644); err != nil {
			return err
		}

		resp, err := http.Get(*user.AvatarURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		file, err := os.Create(filepath.Join(dir, "1.jpg"))
		if err != nil {
			return err
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}
		file.Close()

	}
	return nil
}
