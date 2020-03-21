package models

type ModuleGroup struct {
	ModuleGroupID string   `json:"module_group_id"`
	TDS           float32  `json:"tds"`
	PH            float32  `json:"ph"`
	Humidity      float32  `json:"humidity"`
	LightInterval []int16 `json:"light_interval"`
}