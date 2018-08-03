package scouter

import (
	"log"

	"github.com/globalsign/mgo"
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

func InsertRecord(collection string, record interface{}) error {
	return mongoSession.DB("").C(collection).Insert(record)
}

func UpsertRecord(collection string, id interface{}, record interface{}) (*mgo.ChangeInfo, error) {
	return mongoSession.DB("").C(collection).Upsert(id, record)
}

func UpdateById(collection string, id interface{}, update interface{}) error {
	return mongoSession.DB("").C(collection).UpdateId(id, update)
}
