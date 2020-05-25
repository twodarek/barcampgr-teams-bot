package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"github.com/twodarek/barcampgr-teams-bot/server"
)

func main() {
	router := mux.NewRouter()

	httpClient := &http.Client{}

	teamsClient := webexteams.NewClient()
	initTeamsClient(teamsClient)

	conf := Config{
		APIToken: os.Getenv("CISCO_TEAMS_API_TOKEN"),
	}

	ac := barcampgr.NewAppController(
		teamsClient,
		httpClient,
		conf.APIToken,
	)

	s := server.New(ac, conf.APIToken, router)

	// Multiple codepaths use the DefaultServeMux so we start listening at the top
	go http.ListenAndServe("0.0.0.0:8080", s)

	select {}
}

func initTeamsClient(client *webexteams.Client) error {
	
	return nil
}