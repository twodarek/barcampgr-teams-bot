package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

type AppHandler struct {
	AppController *barcampgr.Controller
	config barcampgr.Config
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
		log.Printf("Error in handling chatop call: %s", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(resultant))
	}
	return
}

func (ah *AppHandler) RootHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello world!"))
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

func (ah *AppHandler) GetSessionJson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["session_str"] != "" {
		session, err := ah.AppController.GetSessionByStr(vars["session_str"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sessionJson, err := json.Marshal(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(sessionJson)
		return
	}
	w.Write([]byte("{}"))
}

func (ah *AppHandler) UpdateSessionJson(w http.ResponseWriter, r *http.Request) {
	//TODO(twodarek) Implement this
	w.Write([]byte("{}"))
}

func (ah *AppHandler) MigrateDatabase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["password"] != "" {
		if vars["password"] == ah.config.AdminPassword {
			err := ah.AppController.MigrateDB()
			if err == nil {
				w.Write([]byte("done"))
				return
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	return
}

func (ah *AppHandler) RollSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["password"] != "" {
		if vars["password"] == ah.config.AdminPassword {
			err := ah.AppController.RollSchedule(vars["sessionBlock"])
			if err == nil {
				w.Write([]byte("done"))
				return
			}
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	return
}

func (ah *AppHandler) GetTimesJson(w http.ResponseWriter, r *http.Request) {
	times, err := ah.AppController.GetTimesJson()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	timesJson, err := json.Marshal(times)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(timesJson)
	return
}

func (ah *AppHandler) GetRoomsJson(w http.ResponseWriter, r *http.Request) {
	rooms, err := ah.AppController.GetRoomsJson()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	roomsJson, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(roomsJson)
	return
}