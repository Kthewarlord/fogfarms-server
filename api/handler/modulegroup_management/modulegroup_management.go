package modulegroup_management

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	r := httprouter.New()
	r.HandlerFunc("GET", "/modulegroup_management", getAllModuleGroup)

	return r
}

func getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}