package scouter

import (
	"github.com/globalsign/mgo/bson"
	"github.com/google/go-github/github"
)

type User struct {
	ID int64 `bson:"_id" json:"id"`
	*github.User
}

func CountUsers() (int, error) {
	return CountCollectionRecords(UserCollection)
}

func FindUsers(selector bson.M, page, pageSize int) ([]User, error) {
	var users []User

	skip := (page - 1) * pageSize
	if skip < 0 {
		skip = 0
	}

	if err := mongoSession.DB("").C(UserCollection).Find(selector).Sort("-$natural").Skip(skip).Limit(pageSize).All(&users); err != nil {
		return users, nil
	}

	return users, nil
}

func InsertUsers(users []github.User) error {
	for _, user := range users {
		u := User{
			ID:   *user.ID,
			User: &user,
		}
		if err := InsertRecord(UserCollection, u); err != nil {
			return err
		}
	}
	return nil
}

func UpsertUser(user github.User) error {
	u := User{
		ID:   *user.ID,
		User: &user,
	}
	_, err := UpsertRecord(UserCollection, bson.M{"_id": u.ID}, u)
	if err != nil {
		return err
	}
	return nil
}
