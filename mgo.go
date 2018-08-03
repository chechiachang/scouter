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

func CountCollectionRecords(collection string) (int, error) {
	return mongoSession.DB("").C(collection).Count()
}

func FindRecord(collection string, query bson.M) (github.User, error) {
	var user github.User
	if err := mongoSession.DB("").C(UserCollection).Find(query).One(&user); err != nil {
		return user, err
	}
	return user, nil
}

func InsertRecord(collection string, record interface{}) error {
	return mongoSession.DB("").C(collection).Insert(record)
}

func UpsertRecord(collection string, id interface{}, record interface{}) (*mgo.ChangeInfo, error) {
	return mongoSession.DB("").C(collection).Upsert(id, record)
}

func FindUser(query bson.M) (github.User, error) {
	return FindRecord(UserCollection, query)
}
