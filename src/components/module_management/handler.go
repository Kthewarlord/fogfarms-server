package module_management

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/module_management/create", createModule).
		Methods("POST").
		Schemes("http")

	return router
}

func createModule(w http.ResponseWriter, r *http.Request) {
	CreateModule(w, r)
}