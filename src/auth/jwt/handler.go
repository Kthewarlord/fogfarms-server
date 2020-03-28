package jwt

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/auth/sign_in", SignIn)

	return router
}

func SignIn(w http.ResponseWriter, r *http.Request) {

}

