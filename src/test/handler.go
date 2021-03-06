package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/gddo/httputil/header"
	"github.com/julienschmidt/httprouter"
)

func MakeHTTPHandler() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("POST", "/test/login", testLogin)
	router.HandlerFunc("POST", "/test/post", PostTestName)
	router.HandlerFunc("GET", "/test/js", GetDemoJson)
	router.HandlerFunc("POST", "/test/arduinoInput", testArduinoInput)
	router.HandlerFunc("GET", "/test/arduinoOutputT", arduinoOutputTrue)
	router.HandlerFunc("GET", "/test/arduinoOutputF", arduinoOutputFalse)
	router.HandlerFunc("POST", "/test/arduinoOutputRand", arduinoOutputRandom)
	rand.Seed(time.Now().UnixNano())
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

func testArduinoInput(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Token          string               `json:"token"`
		SensorData     map[string][]float64 `json:"sensor_data"`
		Mixers         []bool               `json:"mixers"`
		SolenoidValves []bool               `json:"solenoid_valves"`
		LEDs           []bool               `json:"leds"`
		Foggers        []bool               `json:"foggers"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}
	tds := input.SensorData["tds"]
	ph := input.SensorData["ph"]
	solutionTemp := input.SensorData["solution_temp"]
	lux := input.SensorData["grow_unit_lux"]
	humidity := input.SensorData["grow_unit_humidity"]
	temperature := input.SensorData["grow_unit_temp"]
	type InputExtracted struct {
		Input                 Input
		AcceptedNTDS          []float64
		AcceptedNPH           []float64
		AcceptedNSolutionTemo []float64
		AcceptedGLux          []float64
		AcceptedGHumidity     []float64
		AcceptedGTemperature  []float64
	}

	output := InputExtracted{
		Input:        input,
		AcceptedNTDS: tds, AcceptedNPH: ph,
		AcceptedNSolutionTemo: solutionTemp,
		AcceptedGLux:          lux,
		AcceptedGHumidity:     humidity,
		AcceptedGTemperature:  temperature,
	}

	jsonData, err := json.Marshal(output)
	if err != nil {
		msg := "Error: Failed to marshal JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
func arduinoOutputTrue(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		OnAuto bool `json:"on_auto"`
	}

	output := Output{
		OnAuto: true,
	}

	jsonData, err := json.Marshal(output)
	if err != nil {
		msg := "Error: Failed to marshal JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func arduinoOutputFalse(w http.ResponseWriter, r *http.Request) {
	type Output struct {
		OnAuto         bool   `json:"on_auto"`
		Foggers        []bool `json:"foggers"`
		LEDs           []bool `json:"leds"`
		Mixers         []bool `json:"mixers"`
		SolenoidValves []bool `json:"solenoid_valves"`
	}

	output := Output{
		OnAuto:         false,
		Foggers:        []bool{false},
		LEDs:           []bool{false},
		Mixers:         []bool{false},
		SolenoidValves: []bool{false, false},
	}

	jsonData, err := json.Marshal(output)
	if err != nil {
		msg := "Error: Failed to marshal JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func arduinoOutputRandom(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Token          string               `json:"token"`
		SensorData     map[string][]float64 `json:"sensor_data"`
		Mixers         []bool               `json:"mixers"`
		SolenoidValves []bool               `json:"solenoid_valves"`
		LEDs           []bool               `json:"leds"`
		Foggers        []bool               `json:"foggers"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}
	if input.Token != "hihi" {
		msg := "Wrong Token"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	tds := input.SensorData["tds"]
	ph := input.SensorData["ph"]
	solutionTemp := input.SensorData["solution_temp"]
	lux := input.SensorData["grow_unit_lux"]
	humidity := input.SensorData["grow_unit_humidity"]
	temperature := input.SensorData["grow_unit_temp"]

	type InputExtracted struct {
		Input                 Input
		AcceptedNTDS          []float64
		AcceptedNPH           []float64
		AcceptedNSolutionTemo []float64
		AcceptedGLux          []float64
		AcceptedGHumidity     []float64
		AcceptedGTemperature  []float64
	}

	output := InputExtracted{
		Input:        input,
		AcceptedNTDS: tds, AcceptedNPH: ph,
		AcceptedNSolutionTemo: solutionTemp,
		AcceptedGLux:          lux,
		AcceptedGHumidity:     humidity,
		AcceptedGTemperature:  temperature,
	}
	type Output struct {
		Input          InputExtracted `json:"InputData"`
		OnAuto         bool           `json:"on_auto"`
		Foggers        []bool         `json:"foggers"`
		LEDs           []bool         `json:"leds"`
		Mixers         []bool         `json:"mixers"`
		SolenoidValves []bool         `json:"solenoid_valves"`
	}

	var auto bool
	if rand.Float32() < 0.5 {
		auto = true
	} else {
		auto = false
	}
	var output2 Output
	if auto {
		output2 = Output{
			Input:          output,
			OnAuto:         auto,
			Foggers:        nil,
			LEDs:           nil,
			Mixers:         nil,
			SolenoidValves: nil,
		}
	} else {
		output2 = Output{
			Input:          output,
			OnAuto:         auto,
			Foggers:        []bool{rand1()},
			LEDs:           []bool{rand1()},
			Mixers:         []bool{rand1()},
			SolenoidValves: []bool{rand1()},
		}
	}

	jsonData, err := json.Marshal(output2)
	if err != nil {
		msg := "Error: Failed to marshal JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func rand1() bool {
	return rand.Float32() < 0.5
}
