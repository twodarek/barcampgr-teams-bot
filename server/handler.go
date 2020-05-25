package server

import (
	"net/http"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
)

type AppHandler struct {
	AppController *barcampgr.Controller
}

func (ah *AppHandler) GetHelp(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

	resultant, err := ah.AppController.GetHelp(
		//&app.App{
		//	Name: vars["appName"],
		//},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(resultant))
		w.WriteHeader(http.StatusOK)
	}
	return
}