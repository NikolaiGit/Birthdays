package main

import (
	"net/http"
)

var contextpath = "/birthdays"

func muxer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == contextpath {
		getBirthday(w, r)
		return
	}
	if r.URL.Path == contextpath+"/health" {
		healthcheck(w, r)
		return
	}
	if r.URL.Path == contextpath+"/save" {
		saveBirthday(w, r)
		return
	}
	if r.URL.Path == contextpath+"/get" {
		getBirthday(w, r)
		return
	}
	if r.URL.Path == contextpath+"/githubLogin" {
		githubLogin(w, r)
		return
	}
	if r.URL.Path == contextpath+"/githubCallback" {
		githubCallback(w, r)
		return
	}
	if r.URL.Path == contextpath+"/googleLogin" {
		googleLogin(w, r)
		return
	}
	if r.URL.Path == contextpath+"/googleCallback" {
		googleCallback(w, r)
		return
	}
	http.NotFound(w, r)
	return

}

/*type MyMux struct {
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
	if r.URL.Path == contextpath+"/get" {
		getBirthday(w, r)
		return
	}
	if r.URL.Path == contextpath+"/githubLogin" {
		githubLogin(w, r)
		return
	}
	if r.URL.Path == contextpath+"/githubCallback" {
		githubCallback(w, r)
		return
	}
	http.NotFound(w, r)
	return
}*/
