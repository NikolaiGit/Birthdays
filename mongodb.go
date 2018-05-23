package main

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

//MongoDBDialInfo is a object with connection information for mongodb
var MongoDBDialInfo = &mgo.DialInfo{
	//https://godoc.org/labix.org/v2/mgo#DialInfo
	Addrs:    []string{"localhost"},
	Timeout:  3 * time.Second,
	Database: "admin",
	Username: "birthdays-backend",
	Password: "b-b",
}
