package models

import "time"

type ModuleGroup struct {
	ModuleGroupID    int       `json:"module_group_id"`
	ModuleGroupLabel string    `json:"module_group_label"`
	PlantID          int       `json:"plant_id"`
	LocationID       int       `json:"location_id"`
	TDS              float64   `json:"tds"`
	PH               float64   `json:"ph"`
	Humidity         float64   `json:"humidity"`
	OnAuto           bool      `json:"on_auto"`
	LightsOffHour    float64   `json:"lights_off_hour"`
	LightsOnHour     float64   `json:"lights_on_hour"`
	TimerLastReset   time.Time `json:"timer_last_reset"`
}
