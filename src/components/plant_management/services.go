package plant_management

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/plant"
)

func GetAllPlants(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	plants, err := plant.GetAllPlants()
	fmt.Printf("Hi")
	if err != nil {
		msg := "Error: Failed to Get All Plants"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Printf("%+v", plants)
	type Output struct {
		Data []models.Plant
	}
	out := Output{plants}
	plantsJson, err := json.Marshal(out)
	if err != nil {
		msg := "Error: json.Marshal(plants)"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(plantsJson)
}

func CreatePlant(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	input := models.Plant{}

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := plant.CreatePlant(input)
	if err != nil {
		msg := "Error: Failed to New Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func DeletePlant(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	input := models.Plant{}

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := plant.DeletePlant(input.PlantID)
	if err != nil {
		msg := "Error: Failed to Delete Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}


func EditPlant(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	input := &models.Plant{}

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := plant.EditPlant(input)
	if err != nil {
		msg := "Error: Failed to Delete Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}
