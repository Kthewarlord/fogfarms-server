package iot

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/iot/update", update).
		Methods("POST").
		Schemes("http")

	return router
}

func update(w http.ResponseWriter, r *http.Request) {
	Update(w, r)
}