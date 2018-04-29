package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var authServer = "Github"

var (
	jwtTestDefaultKey *rsa.PublicKey
	defaultKeyFunc    jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtTestDefaultKey, nil }
	emptyKeyFunc      jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return nil, nil }
	nilKeyFunc        jwt.Keyfunc
)

//Nicht-Blatt-Handler
func requireTokenAuthentication(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//keine Autorisierung für folgende Pfade nötig
		if r.URL.Path == "/birthdays/githubLogin" || r.URL.Path == "/birthdays/githubCallback" || r.URL.Path == "/birthdays/googleLogin" || r.URL.Path == "/birthdays/googleCallback" {
			inner.ServeHTTP(w, r)
			return
		}

		//check Token vom Request
		//nur für jwt (und die Frage auch, woher er das liest? cookie, session, header?)
		//if token, err := request.ParseFromRequest(r, request.OAuth2Extractor, emptyKeyFunc); err == nil && token.Valid {
		if cookie, err := r.Cookie("token_values"); err == nil {

			//falls vorhanden
			//erzeuge Token aus Informationen in Cookies
			s := cookie.Value
			ss := strings.Split(s, "-")
			token := new(oauth2.Token)
			token.AccessToken = ss[0]
			token.RefreshToken = ss[1]
			token.TokenType = ss[2]
			t := ss[3]
			token.Expiry, _ = time.Parse(time.RFC3339, t)
			if debug {
				fmt.Println("token from cookie: " + token.AccessToken)
			}

			oauthClient := oauthGithubConf.Client(oauth2.NoContext, token)
			client := github.NewClient(oauthClient)
			user, _, err := client.Users.Get(oauth2.NoContext, "")
			if err != nil {
				fmt.Printf("client.Users.Get() failed with '%s'\n", err)
				//TODO check if refreshtoken nutzbar
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			if debug {
				fmt.Println("user: " + user.GetLogin())
				fmt.Println(err)
			}

			//Context Usage
			//https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/
			context := context.WithValue(oauth2.NoContext, ContextKey("username"), ContextValue{user.GetLogin()})
			r = r.WithContext(context)

		} else {

			//falls kein Token vorhanden:
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
