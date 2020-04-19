package models

type Plant struct {
	PlantID       int     `json:"plant_id"`
	Name          string  `json:"name"`
	TDS           float64 `json:"tds"`
	PH            float64 `json:"ph"`
	Lux           float64 `json:"lux"`
	LightsOnHour  float64 `json:"lights_on_hour"`
	LightsOffHour float64 `json:"lights_off_hour"`
}
