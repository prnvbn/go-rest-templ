package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// APIServer adapter around the mux router
type APIServer struct {
	*mux.Router
	cfg *Config
}

// NewServer creates a new server
func NewServer(cfg *Config) *APIServer {
	s := &APIServer{
		Router: &mux.Router{},
		cfg:    cfg,
	}
	s.initRoutes()
	return s
}

func (s APIServer) initRoutes() {
	s.HandleFunc("/", s.homePageHandler).Methods("GET")
	s.HandleFunc("/query", s.nameAsParamHandler).Methods("GET")
	if s.cfg.CatFact.Enabled {
		s.HandleFunc("/catFact", s.catFactHandler).Methods("GET")
	}
	s.HandleFunc("/{name}", s.nameHandler).Methods("GET") // has to be at the bottom
}

func (s APIServer) Run() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	corsHandler := handlers.CORS(originsOk, headersOk, methodsOk)

	addr := fmt.Sprintf("%s:%d", s.cfg.Addr, s.cfg.Port)
	log.Info().Str("addr", addr).Msg("Serving API")
	log.Fatal().Err(http.ListenAndServe(addr, corsHandler(s))).Send()
}
