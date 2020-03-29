package plant_management

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/plant_management", getAllPlants)

	return router
}

func getAllPlants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}