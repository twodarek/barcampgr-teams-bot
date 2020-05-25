package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
)

type Server struct {
	AppController *barcampgr.Controller

	apiToken string

	router *mux.Router
}

func New(
	ac *barcampgr.Controller,
	apiToken string,
	router *mux.Router,
) *Server {
	s := &Server{
		AppController: ac,

		apiToken: apiToken,

		router: router,
	}

	appHandler := AppHandler{
		AppController: ac,
	}


	log.Println("Starting barcampgr-teams-bot")

		// delete routes for primary
	s.router.HandleFunc("/v1/help/{appName}", s.authMiddleWare(appHandler.GetHelp)).Methods("GET")


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
