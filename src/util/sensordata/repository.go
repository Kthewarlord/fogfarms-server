package sensordata

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error)
	RecordSensorData(moduleID int, tds []float64, ph []float64, solutionTemperature []float64,
		lux []float64, humidity []float64, temperature []float64) error
}