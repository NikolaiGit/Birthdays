package main

import (
	"net/http"
)

var contextpath = "/birthdays"

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == contextpath+"/health" {
		healthcheck(w, r)
		return
	}
	if r.URL.Path == contextpath+"/save" {
		saveBirthday(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
