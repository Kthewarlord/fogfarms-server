package plant_management

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	r := httprouter.New()
	r.HandlerFunc("GET", "/plant_management", getAllPlants)

	return r
}

func getAllPlants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}