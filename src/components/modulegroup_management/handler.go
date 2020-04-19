package modulegroup_management

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/modulegroup_management", populateModuleGroupManagementPage).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/create", createModuleGroup).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/assign", assignModuleToModuleGroup).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/edit_modulegroup_label", editModuleGroupLabel).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/change_plant", changePlant).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/modulegroup_management/delete_modulegroup", deleteModuleGroup).
		Methods("POST").
		Schemes("http")

	return router
}

func populateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	PopulateModuleGroupManagementPage(w, r)
}

func createModuleGroup(w http.ResponseWriter, r *http.Request) {
	CreateModuleGroup(w, r)
}

func assignModuleToModuleGroup(w http.ResponseWriter, r *http.Request) {
	AssignModuleToModuleGroup(w, r)
}

func editModuleGroupLabel(w http.ResponseWriter, r *http.Request) {
	EditModuleGroupLabel(w, r)
}

func changePlant(w http.ResponseWriter, r *http.Request) {
	ChangePlant(w, r)
}

func deleteModuleGroup(w http.ResponseWriter, r *http.Request) {
	DeleteModuleGroup(w, r)
}