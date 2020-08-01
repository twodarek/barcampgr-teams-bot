package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
)

type Server struct {
	AppController *barcampgr.Controller
	config barcampgr.Config
	router *mux.Router
}

func New(
	ac *barcampgr.Controller,
	config barcampgr.Config,
	router *mux.Router,
) *Server {
	s := &Server{
		AppController: ac,
		config: config,
		router: router,
	}

	appHandler := AppHandler{
		AppController: ac,
		config: config,
	}


	log.Println("Starting barcampgr-teams-bot")

	// routes for chatops
	s.router.HandleFunc("/api/", s.authMiddleWare(appHandler.RootHello)).Methods("GET")
	s.router.HandleFunc("/api/v1/chatops", s.authMiddleWare(appHandler.HandleChatop)).Methods("POST")
	s.router.HandleFunc("/api/v1/schedule", s.authMiddleWare(appHandler.GetScheduleJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/times", s.authMiddleWare(appHandler.GetTimesJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/rooms", s.authMiddleWare(appHandler.GetRoomsJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/session/{session_str}", s.authMiddleWare(appHandler.GetSessionJson)).Methods("GET")
	s.router.HandleFunc("/api/v1/session/{session_str}", s.authMiddleWare(appHandler.UpdateSessionJson)).Methods("POST")
	//TODO(twodarek) create the below webhook to allow remote creation of the database if need be
	s.router.HandleFunc("/api/v1/migrate/create/{password}", s.authMiddleWare(appHandler.MigrateDatabase)).Methods("GET")
	//TODO(twodarek) create the below webhook to allow any of the organizers to create the next "block" of sessions
	s.router.HandleFunc("/api/v1/migrate/generate/{sessionBlock}/{password}", s.authMiddleWare(appHandler.RollSchedule)).Methods("GET")
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
