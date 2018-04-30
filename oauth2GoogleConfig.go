package main

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

var (
	// Set ClientId and ClientSecret to

	redirectURI = "http://localhost:9090/birthdays/googleCallback"

	// random string for oauth2 API calls to protect against CSRF
	oauthStateStringGoogle = "CSRFBirthdays"
)

func googleLogin(w http.ResponseWriter, r *http.Request) {
	log.Debug("Methode: googleLogin()")
	url := oauthGoogleConf.AuthCodeURL(oauthStateStringGoogle, oauth2.AccessTypeOnline)
	url = url + "&redirect_uri=" + redirectURI
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func googleCallback(w http.ResponseWriter, r *http.Request) {
	log.Debug("Methode: googleCallback()")
	//check CSRF
	state := r.FormValue("state")

	log.Debug("state: " + state)
	if state != oauthStateStringGoogle {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateStringGoogle, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//get token for code
	code := r.FormValue("code")
	log.Debug("code: " + code)
	token, err := oauthGoogleConf.Exchange(oauth2.NoContext, code)
	log.Debug("Token: " + token.TokenType + " " + token.AccessToken)
	if err != nil {
		fmt.Printf("oauthGoogleConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//access google api
	svc, err := calendar.New(oauthGoogleConf.Client(context.Background(), token))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}

	c, err := svc.Colors.Get().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve calendar colors: %v", err)
	}

	log.Printf("Kind of colors: %v", c.Kind)
	log.Printf("Colors last updated: %v", c.Updated)

	for k, v := range c.Calendar {
		log.Printf("Calendar[%v]: Background=%v, Foreground=%v", k, v.Background, v.Foreground)
	}

	for k, v := range c.Event {
		log.Printf("Event[%v]: Background=%v, Foreground=%v", k, v.Background, v.Foreground)
	}

	listRes, err := svc.CalendarList.List().Fields("items/id").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve list of calendars: %v", err)
	}
	for _, v := range listRes.Items {
		log.Printf("Calendar ID: %v\n", v.Id)
	}

	if len(listRes.Items) > 0 {
		id := listRes.Items[0].Id
		res, err := svc.Events.List(id).Fields("items(updated,summary)", "summary", "nextPageToken").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve calendar events list: %v", err)
		}
		for _, v := range res.Items {
			log.Printf("Calendar ID %q event: %v: %q\n", id, v.Updated, v.Summary)
		}
		log.Printf("Calendar ID %q Summary: %v\n", id, res.Summary)
		log.Printf("Calendar ID %q next page token: %v\n", id, res.NextPageToken)
	}
}
