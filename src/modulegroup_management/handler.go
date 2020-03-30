package modulegroup_management

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/modulegroup_management", getDemoJson)
	router.HandlerFunc("GET", "/modulegroup_management/j", getDemoJson)

	return router
}
