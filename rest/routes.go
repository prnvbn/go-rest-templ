package rest

import "net/http"

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}


