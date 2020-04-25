package plant_management

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/plant_management", getAllPlants).
		Methods("GET").
		Schemes("http")

	router.HandleFunc("/plant_management/create_plant", createPlant).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/plant_management/delete_plant", deletePlant).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/plant_management/edit_plant", EditPlant).
		Methods("POST").
		Schemes("http")

	return router
}

func getAllPlants(w http.ResponseWriter, r *http.Request) {
	GetAllPlants(w, r)
}

func createPlant(w http.ResponseWriter, r *http.Request) {
	CreatePlant(w, r)
}

func deletePlant(w http.ResponseWriter, r *http.Request) {
	DeletePlant(w, r)
}

func editPlant(w http.ResponseWriter, r *http.Request) {
	EditPlant(w, r)
}
