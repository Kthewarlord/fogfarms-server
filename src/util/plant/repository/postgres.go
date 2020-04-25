package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetPlantByID(plantID int) (*models.Plant, error) {
	db := database.GetDB()

	sqlStatement := `SELECT * FROM Plant WHERE PlantID = $1`

	rows, err := db.Query(sqlStatement, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plant := &models.Plant{}
	for rows.Next() {

		err := rows.Scan(
			&plant.PlantID,
			&plant.Name,
			&plant.TDS,
			&plant.PH,
			&plant.Lux,
			&plant.LightsOnHour,
			&plant.LightsOffHour,
		)
		if err != nil {
			return nil, err
		}
	}
	return plant, nil
}

func GetAllPlants() ([]models.Plant, error) {
	db := database.GetDB()

	sqlStatement := `SELECT * FROM Plant`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []models.Plant
	for rows.Next() {
		plant := models.Plant{}

		err := rows.Scan(
			&plant.PlantID,
			&plant.Name,
			&plant.TDS,
			&plant.PH,
			&plant.Lux,
			&plant.LightsOffHour,
			&plant.LightsOffHour,
		)
		if err != nil {
			return nil, err
		}

		plants = append(plants, plant)
	}
	return plants, nil
}

func NewPlant(name string, tds float64, ph float64, lux float64, lightsOnHour float64,
	lightsOffHour float64) error {
	db := database.GetDB()

	sqlStatement := `INSERT INTO Plant (Name, TDS, PH, Lux, LightsOnHour, LightsOffHour)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, name, tds, ph, lux, lightsOnHour, lightsOffHour)
	if err != nil {
		return err
	}
	return err
}

func DeletePlant(plantID int) error {
	db := database.GetDB()

	sqlStatement := `DELETE FROM Plant WHERE PlantID = $1`
	_, err := db.Exec(sqlStatement, plantID)
	if err != nil {
		return err
	}
	return err
}

func EditPlant(p *models.Plant) error {
	db := database.GetDB()

	sqlStatement :=
		`UPDATE PLANT
		SET Name = $1, TDS = $2, PH = $3, Lux = $4, LightsOnHour = $5, LightsOffHour = $6
		WHERE PlantID = $7`
	_, err := db.Exec(sqlStatement, p.Name, p.TDS, p.PH, p.Lux, p.LightsOnHour, p.LightsOffHour,
		p.PlantID)
	if err != nil {
		return err
	}
	return err
}
