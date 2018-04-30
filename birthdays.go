package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
	githuboauth "golang.org/x/oauth2/github"
)

//MyLogFormatter ist ein Objekt für das Logging Libary logrus, in welchem eine eigene Timezone gesetzt werden kann
type MyLogFormatter struct {
	log.Formatter
}

//Format ist die zugehörige Methode zu MyLogFormatter
func (u MyLogFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

func init() {

	// INIT Logging //

	//https://github.com/Sirupsen/logrus

	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(MyLogFormatter{&log.TextFormatter{}})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	/* Log Levels
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
	*/
}

//Birthday struct für die Geburtstage
type Birthday struct {
	Name string
	Date string
}

//ContextKey ist nen Type String, der als Key für die Context Key-Value-paare gentuzt wird, da reine string nicht genutzt werden sollen: https://golang.org/pkg/context/#WithValue
type ContextKey string

//ContextValue ist nen Struct als Wrapper für nen String um es in Contexs zu nutzen
type ContextValue struct {
	value string
}

//Get Context Value ist ein Getter für ContextValue
func (v ContextValue) Get() string {
	return v.value
}

//Set Context Value ist ein Getter für ContextValue
func (v ContextValue) Set(value string) {
	v.value = value
}

func main() {
	fmt.Print(githuboauth.Endpoint.AuthURL)
	fmt.Print(githuboauth.Endpoint.TokenURL)
	//http.ListenAndServe(":9090", &MyMux{})
	log.Info("Applikation hört auf :9090")
	//mit oauth-Login-Handler
	//http.Handle("/", alice.New(requireTokenAuthentication).ThenFunc(muxer))
	//ohne oauth-Login-Handler
	http.Handle("/", alice.New(getAuthenticationInformation).ThenFunc(muxer))
	http.ListenAndServe(":9090", nil)

}
