package repository

import (
	"database/sql"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/lib/pq"
)

func GetModuleGroupDevices(moduleGroupID int) ([]models.Device, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT Distinct DeviceID, DeviceTypeID, IsOn, GrowUnitID, NutrientUnitID, PHDownUnitID, PHUpUnitID FROM Device 
		WHERE Device.GrowUnitID IN
		      (SELECT GrowUnit.GrowUnitID FROM GrowUnit
		      WHERE GrowUnit.ModuleID IN
		            (SELECT Module.ModuleID from Module WHERE Module.ModuleGroupID = $1))
		OR Device.NutrientUnitID IN
		   (SELECT NutrientUnit.NutrientUnitID FROM NutrientUnit
		   WHERE NutrientUnit.ModuleID IN
		         (SELECT Module.ModuleID from Module WHERE Module.ModuleGroupID = $1))`

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		type Input struct {
			DeviceID       int
			DeviceTypeID   int
			IsOn           bool
			GrowUnitID     sql.NullInt64
			NutrientUnitID sql.NullInt64
			PHDownUnitID   sql.NullInt64
			PHUpUnitID     sql.NullInt64
		}
		input := Input{}
		var d models.Device

		err = rows.Scan(
			&d.DeviceID,
			&d.DeviceTypeID,
			&d.IsOn,
			&input.GrowUnitID,
			&input.NutrientUnitID,
			&input.PHDownUnitID,
			&input.PHUpUnitID,
		)
		if err != nil {
			return nil, err
		}

		if input.GrowUnitID.Valid {
			d.GrowUnitID = int(input.GrowUnitID.Int64)
		} else {
			d.GrowUnitID = -1
		}

		if input.NutrientUnitID.Valid {
			d.NutrientUnitID = int(input.NutrientUnitID.Int64)
		} else {
			d.NutrientUnitID = -1
		}

		if input.PHDownUnitID.Valid {
			d.PHDownUnitID = int(input.PHDownUnitID.Int64)
		} else {
			d.PHDownUnitID = -1
		}

		if input.PHUpUnitID.Valid {
			d.PHUpUnitID = int(input.PHUpUnitID.Int64)
		} else {
			d.PHDownUnitID = -1
		}

		devices = append(devices, d)
	}

	return devices, nil
}

func ToggleDevice(moduleID int, devicearray []bool, devicetype string) error {
	db := database.GetDB()
	var arr string = ``

	switch dtypes := devicetype; dtypes {
	case "solenoid_valve":
		arr = `arrsolenoidvalve`
	case "mixer":
		arr = `arrmixer`
	case "led":
		arr = `arrled`
	case "fogger":
		arr = `arrfogger`
	default:
		return sql.ErrNoRows
	}

	sqlStatement := `UPDATE Module SET ` + arr + `= $1 WHERE moduleID = $2`

	_, err := db.Query(sqlStatement, pq.Array(devicearray), moduleID)
	return err
}
