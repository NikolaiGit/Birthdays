package main

import (
	"net/http"

	"github.com/justinas/alice"
)

var debug = true

//Birthday struct für die Geburtstage
type Birthday struct {
	Name string
	Date string
}

//ContextKey ist nen Type String, der als Key für die Context Key-Value-paare gentuzt wird, da reine string nicht genutzt werden sollen: https://golang.org/pkg/context/#WithValue
type ContextKey string

//ContextValue ist nen Struct als Wrapper für nen String um es in Contexs zu nutzen
type ContextValue struct {
	value string
}

//Get Context Value ist ein Getter für ContextValue
func (v ContextValue) Get() string {
	return v.value
}

//Set Context Value ist ein Getter für ContextValue
func (v ContextValue) Set(value string) {
	v.value = value
}

func main() {
	//http.ListenAndServe(":9090", &MyMux{})

	http.Handle("/", alice.New(requireTokenAuthentication).ThenFunc(muxer))
	http.ListenAndServe(":9090", nil)
}
