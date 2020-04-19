package models

type NutrientUnit struct {
	NutrientUnitID int `json:"nutrient_unit_id"`
	ModuleID       int `json:"module_id"`
	NutrientID     int `json:"nutrient_id"`
}
