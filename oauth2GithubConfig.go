package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var (
	// Set ClientId and ClientSecret to

	//conf in oauthsecrets.go

	// random string for oauth2 API calls to protect against CSRF
	oauthStateString = "CSRFBirthdays"
)

// /login

//hier redirect auf Callback
//und in callback redirect parameer auf ursprüngliche url
func githubLogin(w http.ResponseWriter, r *http.Request) {
	if debug {
		fmt.Println("githubLogin()")
	}
	url := oauthGithubConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)

	//füge ein URL Parameter hinzu
	//Github leitet nach Erfolgreichem oauth flow an diese Adresse weiter
	//https://developer.github.com/apps/building-oauth-apps/authorization-options-for-oauth-apps/#redirect-urls

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
	url = url + "&redirect_uri=" + "http://localhost:9090/birthdays/githubCallback"
	fmt.Println("neue URL: " + url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// /github_oauth_cb. Called by github after authorization is granted
func githubCallback(w http.ResponseWriter, r *http.Request) {
	if debug {
		fmt.Println("githubCallback()")
	}
	//check CSRF
	state := r.FormValue("state")
	if debug {
		fmt.Println("state: " + state)
	}
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//get token for code
	code := r.FormValue("code")
	if debug {
		fmt.Println("code: " + code)
	}
	token, err := oauthGithubConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthGithubConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if debug {
		fmt.Println("token: " + token.TokenType + " " + token.AccessToken)
	}

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

	//save token in Cookie
	//gute Quelle auch für Attacken: https://www.calhoun.io/securing-cookies-in-go/ !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	c := http.Cookie{
		Name:     "token_values",
		Value:    token.AccessToken + "-" + token.RefreshToken + "-" + token.TokenType + "-" + token.Expiry.Format(time.RFC3339),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &c)

	//r.AddCookie(&c)
	//r.Header.Add("Cookie", `name2="quoted"`)

	//Redirect auf Startseite -> /get
	http.Redirect(w, r, "/birthdays/get", http.StatusTemporaryRedirect)
}
