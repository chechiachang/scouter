package scouter

import (
	"github.com/globalsign/mgo/bson"
)

func UpdateUserContribution(user User, endYear int) error {
	user, err := FindUser(bson.M{"_id": user.ID})
	if err != nil {
		return err
	}

	// Get contribution from created to Now
	startYear := user.CreatedAt.Year()
	//totalContribution := 0
	for ; startYear <= endYear; startYear++ {
		//https://github.com/users/jimmykuo/contributions?from=2016-12-01&to=2016-12-31&full_graph=1

	}

	return UpdateUserById(user.ID, user)
}
