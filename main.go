package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"github.com/twodarek/barcampgr-teams-bot/server"
)

func main() {
	router := mux.NewRouter()

	httpClient := &http.Client{}

	conf := Config{
		APIToken: "",
	}

	ac := barcampgr.NewAppController(
		httpClient,
		conf.APIToken,
	)

	s := server.New(ac, conf.APIToken, router)

	// Multiple codepaths use the DefaultServeMux so we start listening at the top
	go http.ListenAndServe("0.0.0.0:8080", s)

	select {}
}