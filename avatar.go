package scouter

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/google/go-github/github"
)

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

		imagePath := filepath.Join(dir, strconv.Itoa(time.Now().Second())+".jpg")
		file, err := os.Create(imagePath)
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
