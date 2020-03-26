package user_management

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	r := httprouter.New()
	r.HandlerFunc("GET", "/user_management", getAllUsers)

	return r
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}