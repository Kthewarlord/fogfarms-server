package jwt

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/auth/sign_in", SignIn)
	router.HandlerFunc("GET", "/auth/refresh", Refresh)
	router.HandlerFunc("GET", "/auth/sign_out", SignOut)

	return router
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	AuthenticateSignIn(w, r)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	RefreshToken(w, r)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	InvalidateToken(w)
}