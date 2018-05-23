package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var (
	// Set ClientId and ClientSecret to
	//conf in oauthsecrets.go
	// random string for oauth2 API calls to protect against CSRF
	oauthStateString = "CSRFBirthdays"
)

////login
func githubLogin(w http.ResponseWriter, r *http.Request) {
	log.Debug("Methode: githubLogin()")

	url := oauthGithubConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)

	//füge ein URL Parameter hinzu
	//Github leitet nach Erfolgreichem oauth flow an diese Adresse weiter
	//https://developer.github.com/apps/building-oauth-apps/authorization-options-for-oauth-apps/#redirect-urls

	url = url + "&redirect_uri=" + "http://localhost:9090/birthdays/githubCallback"
	log.Info("Redirect auf " + url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	/*
		//oder das hier aus Context auslesen, wenn der redirect von /save auf /githubLogin
		u, err := urlpackage.Parse(r.URL.String())
		if err != nil {
			log.Fatal(err)
		}
		redirectTo := u.Query().Get("redirect_uri")
		if debug {
			fmt.Println("geparste redirect uri auf /githubLogin: " + redirectTo)
		}


	*/
}

// /github_oauth_cb. Called by github after authorization is granted
func githubCallback(w http.ResponseWriter, r *http.Request) {
	log.Debug("Methode: githubCallback()")
	//
	//
	//
	//check CSRF
	state := r.FormValue("state")
	log.Debug("state: " + state)

	if state != oauthStateString {
		log.Error("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//
	//
	//
	//get token for code
	code := r.FormValue("code")
	log.Debug("code: " + code)
	token, err := oauthGithubConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Error("oauthGithubConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.Debug("Token: " + token.TokenType + " " + token.AccessToken)

	//
	//
	//
	//save token in Cookie
	c := http.Cookie{
		Name:     "oauth_token",
		Value:    token.AccessToken + "-" + token.RefreshToken + "-" + token.TokenType + "-" + token.Expiry.Format(time.RFC3339),
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, &c)

	//r.AddCookie(&c)
	//r.Header.Add("Cookie", `name2="quoted"`)
	//TODO
	// save the token in session
	/*
		context := r.Context()
		session := sessions.Default(context)
		session.Set("AccessToken", token.AccessToken)
		session.Set("RefreshToken", token.RefreshToken)
		session.Set("TokenType", token.TokenType)
		session.Set("Expiry", token.Expiry.Format(time.RFC3339))
		session.Save()*/

	//
	//
	//
	//Redirect auf Startseite -> /get
	log.Info("Token von github erhalten und in Cookie token_values gespeichert -> leite nun auf /birthdays/get um")
	http.Redirect(w, r, "/birthdays/get", http.StatusTemporaryRedirect)
}
