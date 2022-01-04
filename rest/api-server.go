package rest

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// APIServer adapter around the mux router
type APIServer struct {
	*mux.Router
}

// NewServer creates a new server
func NewServer() *APIServer {
	s := &APIServer{
		mux.NewRouter(),
	}
	s.initRoutes()
	return s
}

func (s APIServer) initRoutes() {
	s.HandleFunc("/", homePageHandler)
	s.HandleFunc("/query", nameAsQueryHandler)
	s.HandleFunc("/{name}", nameHandler)
}

// Run starts the server
func (s APIServer) Run() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Info("Starting the API server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(s)))
}
