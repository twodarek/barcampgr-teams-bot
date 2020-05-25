package server

import (
	"encoding/json"
	"net/http"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

type AppHandler struct {
	AppController *barcampgr.Controller
}

func (ah *AppHandler) HandleChatop(w http.ResponseWriter, r *http.Request) {
	requestData := webexteams.WebhookRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultant, err := ah.AppController.HandleChatop(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(resultant))
		w.WriteHeader(http.StatusOK)
	}
	return
}


func (ah *AppHandler) GetScheduleJson(w http.ResponseWriter, r *http.Request) {
	schedule, err := ah.AppController.GetScheduleJson()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	scheduleJson, err := json.Marshal(schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(scheduleJson)
	return
}