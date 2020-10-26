package server

import (
	"github.com/twodarek/barcampgr-teams-bot/barcampgr/slack"
	"github.com/twodarek/barcampgr-teams-bot/barcampgr/teams"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
)

type Server struct {
	AppController	   *barcampgr.Controller
	SlackAppController *slack.Controller
	TeamsAppController *teams.Controller
	config barcampgr.Config
	router *mux.Router
}

func New(
	ac *barcampgr.Controller,
	sac *slack.Controller,
	tac *teams.Controller,
	config barcampgr.Config,
	router *mux.Router,
) *Server {
	s := &Server{
		AppController: ac,
		SlackAppController: sac,
		TeamsAppController: tac,
		config: config,
		router: router,
	}

	appHandler := AppHandler{
		AppController: ac,
		SlackAppController: sac,
		TeamsAppController: tac,
		config:             config,
	}


	log.Println("Starting barcampgr-teams-bot")

	s.router.HandleFunc("/api/", s.authMiddleWare(appHandler.RootHello)).Methods("GET")

	// Routes for chatops
	s.router.HandleFunc("/api/v1/slack/chatops", s.authMiddleWare(appHandler.HandleTeamsChatop)).Methods("POST")
	s.router.HandleFunc("/api/v1/webex/chatops", s.authMiddleWare(appHandler.HandleTeamsChatop)).Methods("POST")

	// Routes for membership updates from the main room
	s.router.HandleFunc("/api/v1/slack/membershipUpdates", s.authMiddleWare(appHandler.InviteNewPeopleTeams)).Methods("POST")
	s.router.HandleFunc("/api/v1/webex/membershipUpdates", s.authMiddleWare(appHandler.InviteNewPeopleTeams)).Methods("POST")

	// Routes for web front-end
	s.router.HandleFunc("/api/v1/schedule", s.authMiddleWare(appHandler.GetScheduleJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/times", s.authMiddleWare(appHandler.GetTimesJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/rooms", s.authMiddleWare(appHandler.GetRoomsJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/session/{session_str}", s.authMiddleWare(appHandler.GetSession)).Methods("GET")
	s.router.HandleFunc("/api/v1/session/{session_str}", s.authMiddleWare(appHandler.UpdateSession)).Methods("POST")
	s.router.HandleFunc("/api/v1/session/{session_str}", s.authMiddleWare(appHandler.DeleteSession)).Methods("DELETE")

	// Admin functions
	s.router.HandleFunc("/api/v1/migrate/create/{password}", s.authMiddleWare(appHandler.MigrateDatabase)).Methods("GET")
	s.router.HandleFunc("/api/v1/migrate/generate/{sessionBlock}/{password}", s.authMiddleWare(appHandler.RollSchedule)).Methods("GET")
	//s.router.HandleFunc("/api/v1/webex/invites/new/{password}/", s.authMiddleWare(appHandler.InviteNewEmails)).Methods("POST")

	// Path for static files
	s.router.PathPrefix("/").Handler(http.FileServer(http.Dir("/public/front-end")))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) authMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
	}
}
