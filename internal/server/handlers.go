package server

import (
	"go-rest/pkg/cat"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// nameHandler is a handler for the /{name} endpoint.
// It takes a name as a path parameter and returns a greeting
func (s *Server) nameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	w.Write([]byte("Hello, " + name + "!"))
}

// nameAsParamHandler is a handler for the /query?name={name} endpoint.
// It takes a name as a query parameter and returns a greeting
func (s *Server) nameAsParamHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "name" query parameter from the URL.
	name := r.URL.Query().Get("name")

	w.Write([]byte("Hello, " + name + ", You were passed in as a query!"))
}

func (s *Server) catFactHandler(w http.ResponseWriter, r *http.Request) {
	cr, err := cat.FetchFact(s.cfg.CatFact.URL)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, cr)
}
