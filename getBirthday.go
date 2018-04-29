package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getBirthday(w http.ResponseWriter, r *http.Request) {

	//bessere Context Nutzung : https://medium.com/@matryer/context-keys-in-go-5312346a868d
	username := r.Context().Value(ContextKey("username")).(ContextValue).Get()

	var allBirthdays = getAllBirthdays(username)

	json, err := json.Marshal(allBirthdays)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
	//w.Write([]byte(string(json))
}
