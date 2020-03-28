package modulegroup_management

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/modulegroup_management", getAllModuleGroup)

	return router
}

func getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}