package module_management

import (
	"net/http"

	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
)

func CreateModule(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleLabel string `json:"module_label"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := module.CreateModule(input.ModuleLabel)
	if err != nil {
		msg := "Error: Failed to Create Module"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}
