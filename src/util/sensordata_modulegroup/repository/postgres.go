package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetLatestSensorDataModuleGroup(moduleGroupID int) (*models.SensorDataModuleGroup, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT * FROM SensorData_ModuleGroup WHERE ModuleGroupID = $1
		ORDER BY Timestamp DESC LIMIT 1;`

	rows, err := db.Query(sqlStatement, moduleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensorDataModuleGroup models.SensorDataModuleGroup
	for rows.Next() {
		err := rows.Scan(
			&sensorDataModuleGroup.ModuleGroupID,
			&sensorDataModuleGroup.Timestamp,
			&sensorDataModuleGroup.Humidity,
			&sensorDataModuleGroup.Temperature,
		)
		if err != nil {
			return nil, err
		}
	}

	return &sensorDataModuleGroup, nil
}