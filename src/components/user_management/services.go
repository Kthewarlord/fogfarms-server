package user_management

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"

	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/util/permission"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	u, err := user.GetUserByUsernameFromCookie(r)
	if err != nil {
		msg := "Error: Failed to Get User By Username From Cookie"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if !u.IsAdministrator {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = user.CreateUser(w, r)
	if err != nil {
		msg := "Error: Failed to Create User"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func PopulateUserManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	u, err := user.GetUserByUsernameFromCookie(r)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Request"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	mapUsername, err := permission.PopulateUserManagementPage(u)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Request"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonData, err := json.Marshal(mapUsername)
	if err != nil {
		msg := "Error: Failed to marshal JSON"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func AssignUserModuleGroupPermission(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		Username         string `json:"username"`
		ModuleGroupLabel string `json:"module_group_label"`
		PermissionLevel  int    `json:"permission_level"`
	}

	var input Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}
	fmt.Printf("%+v", input)

	err := permission.AssignUserModuleGroupPermission(input.Username, input.ModuleGroupLabel, input.PermissionLevel)
	if err != nil {
		msg := "Error: Failed to Assign User ModuleGroup Permission"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}
	u, err := user.GetUserByUsernameFromCookie(r)
	if err != nil {
		msg := "Error: Failed to Get User By Username From Cookie"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	if !u.IsAdministrator {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	type Input struct {
		Username string `json:"username"`
	}

	var input Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}
	fmt.Printf("%+v", input)

	err = user.DeleteUserByUsername(input.Username)
	if err != nil {
		msg := "Error: Failed to Delete User By Username"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))

}
