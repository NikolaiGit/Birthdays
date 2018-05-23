package main

import (
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

func getAllBirthdays(kalender string) (results []Birthday) {
	log.Debug("Methode: getAllBirthdays()")
	session, err := mgo.DialWithInfo(MongoDBDialInfo)
	if err != nil {
		log.Error(err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Collection Birthdays
	c := session.DB("birthdays").C("kalender")

	//get all Results for this calendar
	err = c.Find(nil).Sort("-Date").All(&results)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug("Results All: ", results)
	log.Info("Daten aus der Datenbank erhalten")
	return
}
