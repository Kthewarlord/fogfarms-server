package sensordata

import (
	"time"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/sensordata/repository"
)

type Repository interface {
	GetLatestSensorData(moduleGroupID int) ([]models.SensorData, error)
	RecordSensorData(moduleID int, tds []float64, ph []float64, solutionTemperature []float64,
		lux []float64, humidity []float64, temperature []float64) error
}

func GetSensorDataHistory(moduleGroupID int, timeBegin time.Time, timeEnd time.Time) (map[string][]models.SensorData, error) {
	return repository.GetSensorDataHistory(moduleGroupID, timeBegin, timeEnd)
}
