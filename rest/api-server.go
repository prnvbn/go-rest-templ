package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// APIServer adapter around the mux router
type APIServer struct {
	router *mux.Router
}

// NewServer creates a new server
func NewServer() *APIServer {
	s := &APIServer{}
	s.router = mux.NewRouter()
	s.initRoutes()
	return s
}

func (s APIServer) initRoutes() {
	s.router.HandleFunc("/", homePageHandler)
}

// Run starts the server
func (s APIServer) Run() {
	log.Info("Starting server")

	http.ListenAndServe(":3000", s.router)
}
