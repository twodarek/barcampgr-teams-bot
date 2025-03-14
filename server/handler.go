package server

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/slack-go/slack/slackevents"
	bslack "github.com/twodarek/barcampgr-teams-bot/barcampgr/slack"
	"github.com/twodarek/barcampgr-teams-bot/barcampgr/teams"
	"log"
	"net/http"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"

	webexteams "github.com/twodarek/go-cisco-webex-teams/sdk"
)

type AppHandler struct {
	AppController      *barcampgr.Controller
	SlackAppController *bslack.Controller
	TeamsAppController *teams.Controller
	config             barcampgr.Config
}

func (ah *AppHandler) HandleSlackChatop(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: ah.config.SlackVerificationToken}))
	log.Printf("Received: %s", body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error in handling webhook call: %s", err)
		w.Write([]byte(err.Error()))
		return
	}

	resultant := ""
	if eventsAPIEvent.Type == slackevents.URLVerification {
		log.Printf("This is a verification call")
		var challenge *slackevents.ChallengeResponse
		err = json.Unmarshal([]byte(body), &challenge)
		if err == nil {
			w.Header().Set("Content-Type", "text")
			resultant = challenge.Challenge
		}
	} else {
		log.Printf("This is a chatop call")
		resultant, err = ah.SlackAppController.HandleChatop(eventsAPIEvent)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error in handling chatop call: %s", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(resultant))
	}
	return
}

func (ah *AppHandler) HandleTeamsChatop(w http.ResponseWriter, r *http.Request) {
	requestData := webexteams.WebhookRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultant, err := ah.TeamsAppController.HandleChatop(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error in handling chatop call: %s", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(resultant))
	}
	return
}

func (ah *AppHandler) HandleDiscordChatop(w http.ResponseWriter, r *http.Request) {
	hexPubKey, err := hex.DecodeString(ah.config.DiscordPublicKey)
	if err != nil {
		log.Printf("Error decoding hex string pub key: %s", err)
		w.WriteHeader(500)
		return
	}
	pubKey := ed25519.PublicKey(hexPubKey)
	verified := discordgo.VerifyInteraction(r, pubKey)

	log.Printf("verified %v", verified)
	if !verified {
		http.Error(w, "invalid request signature", 401)
		return
	}

	requestData := discordgo.Interaction{}
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received %+v from Discord", requestData)

	if requestData.Type == discordgo.InteractionPing {
		reply := discordgo.Interaction{
			Type: 1,
		}
		err = json.NewEncoder(w).Encode(reply)
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
	//resultant, err := ah.TeamsAppController.HandleChatop(requestData)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	log.Printf("Error in handling chatop call: %s", err)
	//	w.Write([]byte(err.Error()))
	//} else {
	//	w.Write([]byte(resultant))
	//}

	return
}

func (ah *AppHandler) InviteNewPeopleSlack(w http.ResponseWriter, r *http.Request) {
	requestData := webexteams.WebhookRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received membership created webhook, handling: %#v", requestData)
	resultant, err := ah.TeamsAppController.InviteNewPeople(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error in handling chatop call: %s", err)
		w.Write([]byte(err.Error()))
	} else {
		log.Printf("I guess I added people successfully, handled: %#v", requestData)
		w.Write([]byte(resultant))
	}
	return
}

func (ah *AppHandler) InviteNewPeopleTeams(w http.ResponseWriter, r *http.Request) {
	requestData := webexteams.WebhookRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received membership created webhook, handling: %#v", requestData)
	resultant, err := ah.TeamsAppController.InviteNewPeople(requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error in handling chatop call: %s", err)
		w.Write([]byte(err.Error()))
	} else {
		log.Printf("I guess I added people successfully, handled: %#v", requestData)
		w.Write([]byte(resultant))
	}
	return
}

func (ah *AppHandler) InviteNewEmails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["password"] != "" {
		if vars["password"] == ah.config.InvitePassword {
			requestData := barcampgr.InvitePerson{}
			err := json.NewDecoder(r.Body).Decode(&requestData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			log.Printf("Received membership created webhook, handling: %#v", requestData)
			resultant, err := ah.TeamsAppController.InviteNewEmail(requestData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error in handling chatop call: %s", err)
				w.Write([]byte(err.Error()))
			} else {
				log.Printf("I guess I invited the new email successfully, handled: %#v", requestData)
				w.Write([]byte(resultant))
			}
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
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

func (ah *AppHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["session_str"] != "" {
		session, err := ah.AppController.GetSessionByStr(vars["session_str"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if session.UniqueString == "" {
			w.Write([]byte("{}"))
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

func (ah *AppHandler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["session_str"] != "" {
		sessionStr := vars["session_str"]
		requestData := barcampgr.ScheduleSession{}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			log.Printf("Unable to decode the POST from update session %s because %s", sessionStr, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = ah.AppController.UpdateSession(sessionStr, requestData)
		if err != nil {
			log.Printf("Unable to update session %s because %s", sessionStr, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (ah *AppHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["session_str"] != "" {
		sessionStr := vars["session_str"]
		err := ah.AppController.DeleteSession(sessionStr)
		if err != nil {
			log.Printf("Unable to delete session %s because %s", sessionStr, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Session not found", http.StatusNotFound)
	}
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
