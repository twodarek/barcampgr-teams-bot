package main

import (
	"github.com/gorilla/mux"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
)

type Config struct {
	APIToken string
}

func New(
	ac *barcampgr.Controller,
	apiToken string,
	router *mux.Router,
) *Config {
	c := &Config{
		APIToken: apiToken,
	}
	return c
}