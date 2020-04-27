package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetAllLocations() ([]models.Location, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT locationid, city, province FROM Location where locationid !=0;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []models.Location
	for rows.Next() {
		location := models.Location{}

		err := rows.Scan(
			&location.LocationID,
			location.City,
			location.Province,
		)
		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, nil
}
