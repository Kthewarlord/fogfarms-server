package sensordata

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error)
}