package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respondWithError responds with an error message
func respondWithError(w http.ResponseWriter, err error) {
	respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

// respondWithJSON converts the payload into JSON and responds with it
// payload can be struct or a map
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling payload %v", payload)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}
