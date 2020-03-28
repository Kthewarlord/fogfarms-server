package user_management

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/user_management", getAllUsers)

	return router
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}