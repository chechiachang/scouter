package scouter

import (
	"github.com/globalsign/mgo/bson"
	"github.com/google/go-github/github"
)

type User struct {
	ID int64 `bson:"_id" json:"id"`
	*github.User
	Contribution int `bson:"contribution" json:"contribution"`
}

func CountUsers() (int, error) {
	return CountCollectionRecords(UserCollection)
}

func FindUser(query bson.M) (User, error) {
	var user User
	if err := mongoSession.DB("").C(UserCollection).Find(query).One(&user); err != nil {
		return user, err
	}
	return user, nil
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

func InsertUsers(users []User) error {
	for _, user := range users {
		if err := InsertRecord(UserCollection, user); err != nil {
			return err
		}
	}
	return nil
}

func UpdateUserById(id interface{}, update interface{}) error {
	return UpdateById(UserCollection, id, update)
}

func UpsertUser(user User) error {
	_, err := UpsertRecord(UserCollection, bson.M{"_id": user.ID}, user)
	if err != nil {
		return err
	}
	return nil
}

func UpsertUsers(users []User) error {
	for _, user := range users {
		if err := UpsertUser(user); err != nil {
			return err
		}
	}
	return nil
}

func PatchUserContribution(user User) error {
	return UpdateById(UserCollection, bson.M{"_id": user.ID}, bson.M{"contribution": user.Contribution})
}
