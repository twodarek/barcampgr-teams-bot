package barcampgr

import (
	"net/http"
)

type Controller struct {
	httpClient   *http.Client
	apiToken    string
}

func NewAppController(
	httpClient *http.Client,
	apiToken string,

) *Controller {
	return &Controller{
		httpClient:   httpClient,
		apiToken:     apiToken,
	}
}

func (ac *Controller) GetHelp() (string, error) {
	return "This is supposed to be a help text.", nil
}