package rest

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	s.HandleFunc("/", homePageHandler).Methods("GET")
	s.HandleFunc("/query", nameAsParamHandler).Methods("GET")
	s.HandleFunc("/catFact", catFactHandler).Methods("GET")
	s.HandleFunc("/{name}", nameHandler).Methods("GET") // has to be at the bottom
}

// Run starts the server
func (s APIServer) Run() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	corsHandler := handlers.CORS(originsOk, headersOk, methodsOk)
	port, ok := os.LookupEnv("PORT")
	
	if !ok {
		os.Setenv("PORT", "8080")
		logrus.Fatal("$PORT must be set")
	}

	logrus.Info("Starting server on port http://localhost:" + port)
	logrus.Fatal(http.ListenAndServe(":"+port, corsHandler(s)))
}
}
