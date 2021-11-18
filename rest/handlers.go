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

// nameAsQueryHandler is a handler for the /query?name={name} endpoint.
// It takes a name as a query parameter and returns a greeting
func nameAsQueryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	w.Write([]byte("Hello, " + name + ", You were passed in as a query!"))
}
