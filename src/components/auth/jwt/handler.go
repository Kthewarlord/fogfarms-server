package jwt

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/auth/test", AuthenticateTest).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/auth/sign_in", SignIn).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/auth/sign_out", SignOut).
		Methods("GET").
		Schemes("http")
	return router
}

func AuthenticateTest(w http.ResponseWriter, r *http.Request) {
	v := AuthenticateUserToken(w,r)
	log.Println(v)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	AuthenticateSignIn(w, r)
}
func SignOut(w http.ResponseWriter, r *http.Request) {
	InvalidateToken(w)
}
