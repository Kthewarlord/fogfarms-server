package sensordata_modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata_modulegroup/repository"
)

func GetLatestSensorDataModuleGroup(moduleGroupID int) (*models.SensorDataModuleGroup, error) {
	return repository.GetLatestSensorDataModuleGroup(moduleGroupID)
}
