package modulegroup_management

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KitaPDev/fogfarms-server/models"

	"github.com/KitaPDev/fogfarms-server/src/components/auth/jwt"
	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"
	"github.com/KitaPDev/fogfarms-server/src/util/module"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/util/permission"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
)

func PopulateModuleGroupManagementPage(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	u, err := user.GetUserByUsernameFromCookie(r)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Cookie"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var jsonData []byte
	type ModuleGroupData struct {
		ModuleGroupID  int       `json:"module_group_id"`
		PlantID        int       `json:"plant_id"`
		LocationID     int       `json:"location_id"`
		TDS            float64   `json:"tds"`
		PH             float64   `json:"ph"`
		Humidity       float64   `json:"humidity"`
		OnAuto         bool      `json:"on_auto"`
		LightsOffHour  float64   `json:"lights_off_hour"`
		LightsOnHour   float64   `json:"lights_on_hour"`
		TimerLastReset time.Time `json:"timer_last_reset"`
		Permission     int
		Modules        []int
	}
	var moduleGroupMap = make(map[string]*ModuleGroupData)

	var moduleGroupIDs []int
	if u.IsAdministrator {
		moduleGroups, err := modulegroup.GetAllModuleGroups()
		if err != nil {
			msg := "Error: Failed to Get All Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}

		for _, mg := range moduleGroups {
			moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
			moduleGroupMap[mg.ModuleGroupLabel] = &ModuleGroupData{
				ModuleGroupID:  mg.ModuleGroupID,
				PlantID:        mg.PlantID,
				LocationID:     mg.LocationID,
				TDS:            mg.TDS,
				PH:             mg.PH,
				Humidity:       mg.Humidity,
				OnAuto:         mg.OnAuto,
				LightsOffHour:  mg.LightsOffHour,
				LightsOnHour:   mg.LightsOnHour,
				TimerLastReset: mg.TimerLastReset,
				Permission:     4,
				Modules:        []int{},
			}
		}

	} else {
		mapModuleGroupPermissions, err := permission.GetAssignedModuleGroups(u)
		if err != nil {
			msg := "Error: Failed to Get Assigned Module Groups"
			http.Error(w, msg, http.StatusInternalServerError)
			log.Println(err)
			return
		}
		for mg := range mapModuleGroupPermissions {

			moduleGroupMap[mg.ModuleGroupLabel] = &ModuleGroupData{
				ModuleGroupID:  mg.ModuleGroupID,
				PlantID:        mg.PlantID,
				LocationID:     mg.LocationID,
				TDS:            mg.TDS,
				PH:             mg.PH,
				Humidity:       mg.Humidity,
				OnAuto:         mg.OnAuto,
				LightsOffHour:  mg.LightsOffHour,
				LightsOnHour:   mg.LightsOnHour,
				TimerLastReset: mg.TimerLastReset,
				Permission:     mapModuleGroupPermissions[mg],
				Modules:        []int{},
			}
			moduleGroupIDs = append(moduleGroupIDs, mg.ModuleGroupID)
		}

	}

	log.Println("hi", moduleGroupIDs)
	modules, err := module.GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs)
	if err != nil {
		msg := "Error: Failed to Get Modules By ModuleGroupIDs"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	log.Println("hi", modules)
	for _, m := range modules {
		log.Println(m.ModuleLabel)
		log.Println(m.ModuleGroupLabel)
		moduleGroupMap[m.ModuleGroupLabel].Modules = append(moduleGroupMap[m.ModuleGroupLabel].Modules, m.ModuleID)
	}
	jsonData, err = json.Marshal(moduleGroupMap)
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

func CreateModuleGroup(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		PlantID          int     `json:"plant_id"`
		LocationID       int     `json:"location_id"`
		TDS              float64 `json:"tds"`
		PH               float64 `json:"ph"`
		Humidity         float64 `json:"humidity"`
		OnAuto           bool    `json:"on_auto"`
		ModuleGroupLabel string  `json:"module_group_label"`
		LightsOffHour    float64 `json:"lights_off_hour"`
		LightsOnHour     float64 `json:"lights_on_hour"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.CreateModuleGroup(input.ModuleGroupLabel, input.PlantID, input.LocationID,
		input.TDS, input.PH, input.Humidity, input.LightsOnHour, input.LightsOffHour, input.OnAuto)
	if err != nil {
		msg := "Error: Failed to Create Module Group"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func AssignModuleToModuleGroup(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleGroupID int   `json:"module_group_id"`
		ModuleIDs     []int `json:"module_ids"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := module.AssignModulesToModuleGroup(input.ModuleGroupID, input.ModuleIDs)
	if err != nil {
		msg := "Error: Failed to Assign Modules To ModuleGroup"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func EditModuleGroupLabel(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleGroupID    int    `json:"module_group_id"`
		ModuleGroupLabel string `json:"module_group_label"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.EditModuleGroupLabel(input.ModuleGroupID, input.ModuleGroupLabel)
	if err != nil {
		msg := "Error: Failed to Edit ModuleGroup Label"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func ChangePlant(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleGroupID int `json:"module_group_id"`
		PlantID       int `json:"plant_id"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.ChangePlant(input.ModuleGroupID, input.PlantID)
	if err != nil {
		msg := "Error: Failed to Change Plant"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func DeleteModuleGroup(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleGroupID int `json:"module_group_id"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := modulegroup.DeleteModuleGroup(input.ModuleGroupID)
	if err != nil {
		msg := "Error: Failed to Delete ModuleGroup"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func getModuleGroupByMatchedLabel(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		msg := "Unauthorized"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	u, err := user.GetUserByUsernameFromCookie(r)
	if err != nil {
		msg := "Error: Failed to Get User By UserID From Cookie"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	type Input struct {
		ModuleGroupLabel string `json:"module_group_label"`
	}
	var input Input
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}
	var modulegroupQueried []models.ModuleGroup
	if u.IsAdministrator {
		modulegroupQueried, err = modulegroup.GetModuleGroupsByLabelMatch(input.ModuleGroupLabel)
		if err != nil {
			msg := "Error: Failed to Search For ModuleGroup"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	} else {
		modulegroupQueried, err = modulegroup.GetModuleGroupsByLabelMatchForNormal(input.ModuleGroupLabel, u.UserID)
		if err != nil {
			msg := "Error: Failed to Search For ModuleGroup"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	}
	type Output struct {
		Data []models.ModuleGroup
	}
	out := Output{modulegroupQueried}
	jsonData, err := json.Marshal(out)
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

func DeleteModule(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleID int `json:"module_id"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := module.DeleteModule(input.ModuleID)
	if err != nil {
		msg := "Error: Failed to Create Module"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func EditModuleLabel(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleID    int    `json:"module_id"`
		ModuleLabel string `json:"module_label"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	err := module.EditModuleLabel(input.ModuleID, input.ModuleLabel)
	if err != nil {
		msg := "Error: Failed to Edit Module Label"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}
func GetModuleLabel(w http.ResponseWriter, r *http.Request) {
	if !jwt.AuthenticateUserToken(w, r) {
		return
	}

	type Input struct {
		ModuleID int `json:"module_id"`
	}

	input := Input{}
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return
	}

	moduleLabel, err := module.GetModuleLabel(input.ModuleID)
	if err != nil {
		msg := "Error: Failed to Edit Module Label"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	type Output struct {
		ModuleLabel string `json:"module_label"`
	}

	output := Output{ModuleLabel: moduleLabel}

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
