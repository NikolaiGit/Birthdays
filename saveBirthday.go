package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func saveBirthday(w http.ResponseWriter, r *http.Request) {
	fmt.Print("method:", r.Method) //get request method
	fmt.Println(" path", r.URL.Path)
	if r.Method == "GET" {
		saveBirthdayReturnFormHTML(w, r)
	}

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			// Handle error here via logging and then return
		}

		fmt.Println("Name:", r.PostFormValue("name"))
		fmt.Println("Birthday:", r.PostFormValue("birthday"))

		birthdate := Birthday{string(r.PostFormValue("name")), string(r.PostFormValue("birthday"))}

		//Abspeichern
		persistBirthday(birthdate, "test")
		//TODO
		//Username Übergabe für Collection = Kalendername für mehrere Kalender

		//Antwort
		saveBirthdayResponse(w, r, birthdate)
		/**if m, _ := regexp.MatchString("text/html", r.Header.Get("Content-type")); !m {
			fmt.Fprintf(w, "<html><head><title>Accepted!</title></head><body><p>Geburtsgag von %v am %v abgespeichert!</p></body></html>", birthdate.Name, birthdate.Date)
		}*/

	}
}
func saveBirthdayValidateInput(w http.ResponseWriter, r *http.Request, b Birthday) {
	if len(r.Form["name"][0]) == 0 {
		panic("Keine Eingabe für Name")
	}
	if len(r.Form["birthday"][0]) == 0 {
		panic("Keine Eingabe für Birthday")
	}
}

func saveBirthdayReturnFormHTML(w http.ResponseWriter, r *http.Request) {
	gopath := os.Getenv("GOPATH")
	t, err := template.ParseFiles(filepath.Join(gopath, "/src/birthdays/resources/enterBirthday.gtpl"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Execute(w, nil))
}

func saveBirthdayResponse(w http.ResponseWriter, r *http.Request, b Birthday) {
	//if m, _ := regexp.MatchString("application/json", r.Header.Get("Content-type")); !m {
	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(json))
}
