package models

type Device struct {
	DeviceID       int  `json:"device_id"`
	DeviceTypeID   int  `json:"device_type_id"`
	IsOn           bool `json:"IsOn"`
	GrowUnitID     int  `json:"grow_unit_id"`
	NutrientUnitID int  `json:"nutrient_unit_id"`
	PHDownUnitID   int  `json:"ph_down_unit_id"`
	PHUpUnitID     int  `json:"ph_up_unit_id"`
}
