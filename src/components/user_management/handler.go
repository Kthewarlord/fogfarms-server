package user_management

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/user_management", populateUserManagementPage).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/user_management/register", register).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/user_management/assign", assignUserModuleGroupPermission).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/user_management/delete", deleteUser).
		Methods("POST").
		Schemes("http")

	return router
}

func register(w http.ResponseWriter, r *http.Request) {
	CreateUser(w, r)
}

func populateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	PopulateUserManagementPage(w, r)
}

func assignUserModuleGroupPermission(w http.ResponseWriter, r *http.Request) {
	AssignUserModuleGroupPermission(w, r)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	DeleteUser(w, r)
}
