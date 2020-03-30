package modulegroup_management

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/modulegroup_management", getAllModuleGroup)

	return router
}
