package repository

import (
	"log"
	"strings"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/lib/pq"
)

func GetAllModuleGroups() ([]models.ModuleGroup, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, Param_TDs, Param_PH, 
		Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset FROM ModuleGroup;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moduleGroups []models.ModuleGroup
	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			return nil, err
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	return moduleGroups, nil
}

func GetModuleGroupByID(moduleGroupID int) (*models.ModuleGroup, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, Param_TDs, Param_PH, 
		Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset
		FROM ModuleGroup WHERE ModuleGroupID = $1;`

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	moduleGroup := &models.ModuleGroup{}
	for rows.Next() {

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			return nil, err
		}
	}

	return moduleGroup, nil
}

func GetModuleGroupsByIDs(moduleGroupIDs []int) ([]models.ModuleGroup, error) {
	var moduleGroups []models.ModuleGroup
	var err error

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, Param_TDs, Param_PH, 
		Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset
		FROM ModuleGroup WHERE ModuleGroupID = ANY($1);`

	db := database.GetDB()

	rows, err := db.Query(sqlStatement, pq.Array(moduleGroupIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.LocationID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.OnAuto,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			return nil, err
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	log.Println("Variable moduleGroups in GetModuleGroups by ID", moduleGroups)

	return moduleGroups, err
}

func CreateModuleGroup(label string, plantID int, locationID int, tds float64, ph float64,
	humidity float64, lightsOn float64, lightsOff float64, onAuto bool) error {
	db := database.GetDB()

	sqlStatement :=
		`INSERT INTO ModuleGroup (ModuleGroupLabel, PlantID, LocationID, onAuto,
		 Param_TDS, Param_Ph, Param_Humidity, LightsOnHour, LightsOffHour, TimerLastReset)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, Now())`
	_, err := db.Exec(sqlStatement, label, plantID, locationID, onAuto, tds, ph, humidity,
		lightsOn, lightsOff)
	if err != nil {
		return err
	}

	return nil
}

func ToggleAuto(moduleGroupID int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET OnAuto = NOT OnAuto WHERE ModuleGroupID = $1`
	_, err := db.Exec(sqlStatement, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func SetEnvironmentParameters(moduleGroupID int, humidity float64, ph float64, tds float64,
	lightsOnHour float64, lightsOffHour float64) error {
	db := database.GetDB()

	sqlStatement :=
		`UPDATE ModuleGroup	
			SET Param_Humidity = $1, Param_PH = $2, Param_TDS = $3, LightsOnHour = $4, 
				LightsOffHour = $5
			WHERE ModuleGroupID = $6`
	_, err := db.Exec(sqlStatement, humidity, ph, tds, lightsOffHour, lightsOnHour, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func EditModuleGroupLabel(moduleGroupID int, moduleGroupLabel string) error {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET ModuleGroupLabel = $1 WHERE ModuleGroupID = $2`

	_, err := db.Exec(sqlStatement, moduleGroupLabel, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteModuleGroup(moduleGroupID int) error {
	db := database.GetDB()

	sqlStatement := `DELETE FROM ModuleGroup WHERE ModuleGroupID = $1`

	_, err := db.Exec(sqlStatement, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func ChangePlant(moduleGroupID int, plantID int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET PlantID = $1 WHERE ModuleGroupID = $2`

	_, err := db.Exec(sqlStatement, plantID, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func ResetTimer(moduleGroupID int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE ModuleGroup SET TimerLastReset = NOW() WHERE ModuleGroupID = $1;`

	_, err := db.Query(sqlStatement, moduleGroupID)

	return err
}

func GetOnAutoByModuleID(moduleID int) (bool, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT onAuto FROM ModuleGroup
		WHERE ModuleGroupID = (SELECT ModuleGroupID FROM Module WHERE ModuleID = $1)`

	rows, err := db.Query(sqlStatement, moduleID)
	if err != nil {
		return false, err
	}

	var onAuto bool
	for rows.Next() {
		err = rows.Scan(&onAuto)

		if err != nil {
			return false, nil
		}
	}

	return onAuto, nil
}

func GetModuleGroupsByLabelMatch(moduleGroupLabel string) ([]models.ModuleGroup, error) {
	var moduleGroups []models.ModuleGroup
	var err error

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, locationid,Param_TDs, Param_PH, 
		Param_Humidity, onauto,LightsOnHour, LightsOffHour, TimerLastReset
		FROM ModuleGroup WHERE Lower(ModuleGroupLabel) LIKE $1||'%' ;`
	db := database.GetDB()
	rows, err := db.Query(sqlStatement, strings.ToLower(moduleGroupLabel))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.LocationID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.OnAuto,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			log.Println(err)
			return nil, err

		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	log.Println("Variable moduleGroups in GetModuleGroups by ID", moduleGroups)

	return moduleGroups, err
}

func GetModuleGroupsByLabelMatchForNormal(moduleGroupLabel string, userID int) ([]models.ModuleGroup, error) {
	var moduleGroups []models.ModuleGroup
	var err error

	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, locationid,Param_TDs, Param_PH, 
		Param_Humidity, onauto,LightsOnHour, LightsOffHour, TimerLastReset
		FROM ModuleGroup WHERE Lower(ModuleGroupLabel) LIKE $1||'%' AND modulegroupid IN (SELECT modulegroupid from permission where userid=$2);`
	db := database.GetDB()
	rows, err := db.Query(sqlStatement, strings.ToLower(moduleGroupLabel), userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.LocationID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.OnAuto,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			log.Println(err)
			return nil, err

		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	log.Println("Variable moduleGroups in GetModuleGroups by ID", moduleGroups)

	return moduleGroups, err
}

func GetModuleGroupByModuleGroupID(moduleGroupID int) ([]models.ModuleGroup, error) {
	sqlStatement :=
		`SELECT ModuleGroupID, ModuleGroupLabel, PlantID, locationid,Param_TDs, Param_PH, 
	Param_Humidity, onauto,LightsOnHour, LightsOffHour, TimerLastReset
	FROM ModuleGroup WHERE modulegroupid=$1 ;`
	db := database.GetDB()
	var moduleGroups []models.ModuleGroup

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.LocationID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.OnAuto,
			&moduleGroup.LightsOnHour,
			&moduleGroup.LightsOffHour,
			&moduleGroup.TimerLastReset,
		)
		if err != nil {
			log.Println(err)
			return nil, err

		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	log.Println("Variable moduleGroups in GetModuleGroups by ID", moduleGroups)

	return moduleGroups, err

}
