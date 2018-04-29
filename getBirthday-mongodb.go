package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

func getAllBirthdays(kalender string) (results []Birthday) {
	session, err := mgo.DialWithInfo(MongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Collection Birthdays
	c := session.DB("birthdays").C("kalender")

	//get all Results for this calendar
	err = c.Find(nil).Sort("-Date").All(&results)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Results All: ", results)
		return
	}
}
