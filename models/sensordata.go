package models

import "time"

type SensorData struct {
	TimeStamp           time.Time `json:"timestamp"`
	ModuleID            string    `json:"module_id"`
	TDS                 float32   `json:"tds"`
	PH                  float32   `json:"ph"`
	SolutionTemperature float32   `json:"soln_temp"`
	GrowUnitLux         []float32 `json:"grow_unit_lux"`
	GrowUnitHumidity    []float32 `json:"grow_unit_humidity"`
	GrowUnitTemperature []float32 `json:"grow_unit_temp"`
	Humidity            bool      `json:"humidity"`
	Temperature         float32   `json:"temp"`
}