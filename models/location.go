package models

type Location struct {
	LocationID string `json:"location_id"`
	City       string `json:"city"`
	Province   string `json:"province"`
}
