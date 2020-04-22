package location

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/location/repository"
)

func GetAllLocations() ([]models.Location, error) {
	Locations, err := repository.GetAllLocations()
	return Locations, err
}
