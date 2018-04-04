package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func saveBirthday(w http.ResponseWriter, r *http.Request) {
	fmt.Print("method:", r.Method) //get request method
	fmt.Println(" path", r.URL.Path)
	if r.Method == "GET" {
		gopath := os.Getenv("GOPATH")
		t, err := template.ParseFiles(filepath.Join(gopath, "/src/birthdays/resources/enterBirthday.gtpl"))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(t.Execute(w, nil))
	}

	if r.Method == "POST" {

		r.ParseForm()
		for i := range r.Form {
			fmt.Println(i)
		}
		fmt.Println("Name:", r.Form["name"])
		fmt.Println("Birthday:", r.Form["birthday"])

		if len(r.Form["name"][0]) == 0 {
			panic("Keine Eingabe für Name")
		}
		if len(r.Form["birthday"][0]) == 0 {
			panic("Keine Eingabe für Birthday")
		}

		birthdate := Birthday{string(r.Form["name"][0]), string(r.Form["birthday"][0])}

		//Abspeichern

		//Antwort
		if m, _ := regexp.MatchString("text/html", r.Header.Get("Content-type")); !m {
			fmt.Fprintf(w, "<html><head><title>Accepted!</title></head><body><p>Geburtsgag von %v am %v abgespeichert!</p></body></html>", birthdate.Name, birthdate.Date)
		}

		if m, _ := regexp.MatchString("application/json", r.Header.Get("Content-type")); !m {
			json, err := json.Marshal(birthdate)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(w, string(json))
		}
	}
}
