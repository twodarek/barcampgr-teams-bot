package barcampgr

import (
	"errors"
	"fmt"
	"log"
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

func (ac *Controller) HandleChatop(requestData webexteams.WebhookRequest) (string, error) {
	message, _, err := ac.teamsClient.Messages.GetMessage(requestData.Data.ID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to get message id %s", requestData.Data.ID))
	}
	log.Printf("Received message `%s` as message id %s", message.Text, requestData.Data.ID)
	//TODO(twodarek): Figure out what the message sent was and deal with it
	return "", nil
}

func (ac *Controller) GetScheduleJson() (Schedule, error) {
	schedule := Schedule{
		RefreshedAt: "",
		LastUpdate:  "",
		Times:       nil,
		Rows:        nil,
	}
	return schedule, nil
}