package models

import "time"

type SensorDataModuleGroup struct {
	ModuleGroupID int       `json:"module_group_id"`
	Timestamp     time.Time `json:"timestamp"`
	Humidity      float64   `json:"humidity"`
	Temperature   float64   `json:"temperature"`
}
