package main

import mgo "gopkg.in/mgo.v2"

func persistBirthday(b Birthday, kalender string) {

	session, err := mgo.DialWithInfo(MongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Collection Birthdays
	c := session.DB("birthdays").C("kalender")

	err = c.Insert(&b)
	if err != nil {
		panic(err)
	}

}
