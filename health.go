package main

import (
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	fmt.Println("path", r.URL.Path)

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, string(http.StatusServiceUnavailable))
	}
	session.Close()

	fmt.Fprintf(w, "Mir gehts gut!") // send data to client side
}
