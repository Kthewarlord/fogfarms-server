package sensordata_modulegroup

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetLatestSensorDataModuleGroup(moduleGroupID int) (*models.SensorDataModuleGroup, error)
}