package main

import (
	"net/http"
)

type Birthday struct {
	Name string
	Date string
}

func main() {
	http.ListenAndServe(":9090", &MyMux{})
}
