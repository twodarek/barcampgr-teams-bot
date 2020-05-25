package barcampgr

import (
	"net/http"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

type Controller struct {
	teamsClient	*webexteams.Client
	httpClient   *http.Client
	apiToken    string
}

func NewAppController(
	teamsClient	*webexteams.Client,
	httpClient *http.Client,
	apiToken string,

) *Controller {
	return &Controller{
		teamsClient:  teamsClient,
		httpClient:   httpClient,
		apiToken:     apiToken,
	}
}

func (ac *Controller) GetHelp() (string, error) {
	return "This is supposed to be a help text.", nil
}