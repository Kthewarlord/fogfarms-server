package models

type Plant struct {
	PlantID string  `json:"plant_id"`
	Name    string  `json:"name"`
	TDS     float32 `json:"tds"`
	PH      float32 `json:"ph"`
	Lux     float32 `json:"lux"`
}
