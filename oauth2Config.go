package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var authServer = "Github"

//Nicht-Blatt-Handler
func requireTokenAuthentication(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info(r.Method + "-Request auf " + r.URL.Path + " von " + r.RemoteAddr)

		//keine Autorisierung für folgende Pfade nötig
		if r.URL.Path == "/birthdays/githubLogin" || r.URL.Path == "/birthdays/githubCallback" || r.URL.Path == "/birthdays/googleLogin" || r.URL.Path == "/birthdays/googleCallback" {
			inner.ServeHTTP(w, r)
			return
		}

		//check Token vom Request

		//nur für jwt (und die Frage auch, woher er das liest? cookie, session, header?)
		//if token, err := request.ParseFromRequest(r, request.OAuth2Extractor, emptyKeyFunc); err == nil && token.Valid {

		log.Info("prüfe auf Token in Cookie")
		if cookie, err := r.Cookie("token_values"); err == nil {

			//falls vorhanden
			log.Info("Token gefunden")
			//erzeuge Token aus Informationen in Cookies
			s := cookie.Value
			ss := strings.Split(s, "-")
			token := new(oauth2.Token)
			token.AccessToken = ss[0]
			token.RefreshToken = ss[1]
			token.TokenType = ss[2]
			t := ss[3]
			token.Expiry, _ = time.Parse(time.RFC3339, t)
			log.Debug("Token from Cookie: " + token.AccessToken)

			oauthClient := oauthGithubConf.Client(oauth2.NoContext, token)
			client := github.NewClient(oauthClient)
			user, _, err := client.Users.Get(oauth2.NoContext, "")
			if err != nil {
				fmt.Printf("client.Users.Get() failed with '%s'\n", err)
				//TODO check if refreshtoken nutzbar
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			log.Debug("User: " + user.GetLogin())

			//Context Usage
			//https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/
			context := context.WithValue(oauth2.NoContext, ContextKey("username"), ContextValue{user.GetLogin()})
			r = r.WithContext(context)

		} else {

			//falls kein Token vorhanden:
			log.Info("Kein Token gefunden -> leite auf OAuth-Autorisierung um")
			redirectToLogin(w, r)
			//fmt.Println("Authentication failed " + err.Error())
			//w.WriteHeader(http.StatusForbidden)
			return
		}

		inner.ServeHTTP(w, r)
	})
	//brauch ich nicht, da obe in zweiter Zeile returned wird
	//return http.HandlerFunc(fn)
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	//redirectUri := "?redirect_uri=http://localhost:9090" + r.URL.Path

	switch authServer {
	case "Google":
		//http.Redirect(w, r, "/birthdays/googleLogin"+redirectUri, http.StatusTemporaryRedirect)
		http.Redirect(w, r, "/birthdays/googleLogin", http.StatusTemporaryRedirect)

	case "Github":
		//http.Redirect(w, r, "/birthdays/githubLogin"+redirectUri, http.StatusTemporaryRedirect)
		http.Redirect(w, r, "/birthdays/githubLogin", http.StatusTemporaryRedirect)

	}
}
