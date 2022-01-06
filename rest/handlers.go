package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// nameHandler is a handler for the /{name} endpoint.
// It takes a name as a path parameter and returns a greeting
func nameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	w.Write([]byte("Hello, " + name + "!"))
}

// nameAsParamHandler is a handler for the /query?name={name} endpoint.
// It takes a name as a query parameter and returns a greeting
func nameAsParamHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Write([]byte("Hello, " + name + ", You were passed in as a query!"))
}

func catFactHandler(w http.ResponseWriter, r *http.Request) {
	cr, err := getCatFact()
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, cr)
}
