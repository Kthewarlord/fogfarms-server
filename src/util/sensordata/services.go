package sensordata

import (
	"github.com/KitaPDev/fogfarms-server/models/outputs"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata/repository"
)

func GetLatestSensorData(moduleGroupID int) (map[string]*outputs.Dashboardoutput, error) {
	return repository.GetLatestSensorData(moduleGroupID)
}

func RecordSensorData(moduleID int, tds []float64, ph []float64, solutionTemperature []float64,
	lux []float64, humidity []float64, temperature []float64) error {

	return repository.RecordSensorData(moduleID, tds, ph, solutionTemperature, lux, humidity, temperature)
}