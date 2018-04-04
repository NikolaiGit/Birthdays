package main

import (
	"fmt"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "Mir gehts gut!") // send data to client side
}
