package plant

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetPlant(plantID string) *models.Plant
	GetAllPlants() []models.Plant
	NewPlant(name string, tds float64, ph float64, lux float64)
	DeletePlant(plantID int) error
	EditPlant(p *models.Plant) error
}