package models

type Location struct {
	LocationID int    `json:"location_id"`
	City       string `json:"city"`
	Province   string `json:"province"`
}
