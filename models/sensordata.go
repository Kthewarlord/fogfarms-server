package models

import "time"

type SensorData struct {
	ModuleID            int       `json:"module_id"`
	TimeStamp           time.Time `json:"timestamp"`
	TDS                 []float64 `json:"tds"`
	PH                  []float64 `json:"ph"`
	SolutionTemperature []float64 `json:"solution_temp"`
	GrowUnitLux         []float64 `json:"grow_unit_lux"`
	GrowUnitHumidity    []float64 `json:"grow_unit_humidity"`
	GrowUnitTemperature []float64 `json:"grow_unit_temp"`
}
