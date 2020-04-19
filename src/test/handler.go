package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/gddo/httputil/header"
	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("POST", "/test/login", testLogin)
	router.HandlerFunc("POST", "/test/post", PostTestName)
	router.HandlerFunc("GET", "/test/js", GetDemoJson)

	return router
}

func GetDemoJson(w http.ResponseWriter, r *http.Request) {
	type DemoJson struct {
		Name string
		Item []string
	}
	demoJson := DemoJson{"Name", []string{"item 1", "item 2"}}
	js, err := json.Marshal(demoJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func PostTestName(w http.ResponseWriter, r *http.Request) {
	type TestData struct {
		Name string
	}
	var testData TestData
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	err := json.NewDecoder(r.Body).Decode(&testData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Print("%+v", testData)
	js, err := json.Marshal(testData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func testLogin(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username string
		Password string
	}
	var testdata Input
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	err := json.NewDecoder(r.Body).Decode(&testdata)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Printf("%+v", testdata)
	username := testdata.Username
	password := testdata.Password
	if username == "ddfsdd" && password == "hihi" {
		GenerateToken(username, w)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func GenerateToken(username string, w http.ResponseWriter) {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	expirationTime := time.Now().Add(10 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  expirationTime,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(io.WriteString(w, `{"error":"token_generation_failed"`))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwtToken",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
