package main

import (
	"github.com/chechiachang/scouter"
	"github.com/globalsign/mgo/bson"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {

	total, err := scouter.CountUsers()
	if err != nil {
		log.Fatal(err)
	}

	pageNum := total / scouter.SearchMaxPerPage

	query := bson.M{}
	sort := "$natural"

	for page := 1; page < pageNum+1; page++ {

		users, err := scouter.FindUsers(query, sort, page, scouter.SearchMaxPerPage)
		if err != nil {
			log.Fatal(err)
		}

		GetAvatar(users)
	}
}

func GetAvatar(users []scouter.User) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Println("work dir: " + wd)
	workspace := filepath.Join(wd, "data/avatars")
	if _, err := os.Stat(workspace); os.IsNotExist(err) {
		log.Println(workspace + " not found. Mkdir one.")
		if err = os.MkdirAll(workspace, 0755); err != nil {
			return err
		}
	}

	wg := sync.WaitGroup{}
	defer wg.Done()

	for _, user := range users {
		wg.Add(1)
		defer wg.Done()

		go downloadAvatar(workspace, user)

	}
	wg.Wait()
	return nil
}

func downloadAvatar(workspace string, user scouter.User) error {
	resp, err := http.Get(*user.AvatarURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imagePath := filepath.Join(workspace, strconv.FormatInt(user.ID, 10)+".jpg")
	log.Println("Downloading image: ", imagePath)
	file, err := os.Create(imagePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
