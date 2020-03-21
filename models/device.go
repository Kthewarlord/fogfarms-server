package models

type Device struct {
	DeviceID string `json:"device_id"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}