package jsonhandler

import (
	"encoding/json"
	"github.com/golang/gddo/httputil/header"
	"log"
	"net/http"
)

func DecodeJsonFromBody(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return false
		}
	}
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		msg := "Error: Failed to Decode JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return false
	}

	return true
}
