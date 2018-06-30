package scouter

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/go-github/github"
)

const (
	MongoUrl       = "mongodb://127.0.0.1:27017/scouter"
	UserCollection = "users"
)

var mongoSession *mgo.Session

func init() {
	session, err := mgo.Dial(MongoUrl)
	if err != nil {
		log.Fatal(err)
	}
	mongoSession = session
}

func FindUser() *github.User {
	c := mongoSession.DB("").C(UserCollection)

	result := new(github.User)

	query := bson.M{}
	err := c.Find(query).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func InsertUsers(users []github.User) error {
	c := mongoSession.DB("").C(UserCollection)

	for _, user := range users {
		u := User{
			ID:   *user.ID,
			User: &user,
		}
		err := c.Insert(u)
		if err != nil {
			return err
		}
	}
	return nil
}
