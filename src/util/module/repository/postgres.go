package repository

import (
	"log"
	"math/rand"
	"time"

	"github.com/KitaPDev/fogfarms-server/models/outputs"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/lib/pq"
)

func CreateModule(moduleLabel string) error {
	db := database.GetDB()

	sqlStatement :=
		`INSERT INTO Module (ModuleLabel, Token, ArrFogger, ArrLED, ArrMixer, ArrSolenoidValve)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Query(sqlStatement, moduleLabel, GenerateToken(), pq.BoolArray{}, pq.BoolArray{},
		pq.BoolArray{}, pq.BoolArray{})
	if err != nil {
		return err
	}

	return nil
}

func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ModuleID, ModuleGroupID, ModuleLabel FROM Module WHERE ModuleGroupID = ANY($1)`

	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []models.Module
	for rows.Next() {
		module := models.Module{}

		err := rows.Scan(&module.ModuleID, &module.ModuleGroupID, &module.ModuleLabel)
		if err != nil {
			return nil, err
		}

		modules = append(modules, module)
	}
	return modules, nil
}

func GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs []int) ([]outputs.ModuleOutput, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT module.moduleID, module.moduleGroupID, modulelabel,modulegrouplabel FROM Module,Modulegroup WHERE module.ModuleGroupID = ANY($1) AND module.modulegroupID=modulegroup.modulegroupID ;`

	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modules []outputs.ModuleOutput
	for rows.Next() {
		module := outputs.ModuleOutput{}
		log.Println(module)
		err := rows.Scan(&module.ModuleID, &module.ModuleGroupID, &module.ModuleLabel, &module.ModuleGroupLabel)
		if err != nil {
			return nil, err
		}
		log.Println(module)
		modules = append(modules, module)
	}
	return modules, nil
}

func AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE Module SET ModuleGroupID = $1 WHERE ModuleID = ANY($2)`

	_, err := db.Query(sqlStatement, moduleGroupID, pq.Array(moduleIDs))

	return err
}

func DeleteModule(moduleLabel string) error {
	db := database.GetDB()

	sqlStatement := `DELETE FROM Module WHERE ModuleLabel = $1;`

	_, err := db.Query(sqlStatement, moduleLabel)
	if err != nil {
		return err
	}

	return nil
}

func EditModuleLabel(moduleID int, moduleLabel string) error {
	db := database.GetDB()

	sqlStatement := `UPDATE Module SET ModuleLabel = $1 WHERE ModuleID = $2`

	_, err := db.Query(sqlStatement, moduleLabel, moduleID)
	if err != nil {
		return err
	}

	return nil
}

func GetModuleIDByToken(token string) (int, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ModuleID FROM Module WHERE Token = $1`

	rows, err := db.Query(sqlStatement, token)
	if err != nil {
		return -1, err
	}

	var moduleID int
	for rows.Next() {
		err := rows.Scan(&moduleID)
		if err != nil {
			return -1, err
		}
	}

	return moduleID, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetDeviceStatus(moduleID int) ([]bool, []bool, []bool, []bool, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ArrFogger, ArrLED, ArrMixer, ArrSolenoidValve FROM Module
		WHERE ModuleID = $1;`

	rows, err := db.Query(sqlStatement, moduleID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var fogger, led, mixer, solenoidValve []bool
	for rows.Next() {
		err = rows.Scan(
			pq.Array(&fogger),
			pq.Array(&led),
			pq.Array(&mixer),
			pq.Array(&solenoidValve),
		)

		if err != nil {
			return nil, nil, nil, nil, err
		}
	}

	return fogger, led, mixer, solenoidValve, nil
}

func UpdateDeviceStatus(moduleID int, mixer []bool, solenoidValves []bool, led []bool, fogger []bool) error {
	db := database.GetDB()

	sqlStatement :=
		`UPDATE Module SET ArrMixer = $1, ArrSolenoidValve = $2, ArrLED = $3, ArrFogger = $4
		WHERE ModuleID = $5;`

	_, err := db.Query(sqlStatement, pq.Array(mixer), pq.Array(solenoidValves), pq.Array(led),
		pq.Array(fogger), moduleID)
	if err != nil {
		return err
	}

	return nil
}
