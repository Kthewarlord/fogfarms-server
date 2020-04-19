package sensordata

import (
	"github.com/KitaPDev/fogfarms-server/models/outputs"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata/repository"
)

func GetLatestSensorData(moduleGroupID int) (map[string]*outputs.Dashboardoutput, error) {
	return repository.GetLatestSensorData(moduleGroupID)
}
