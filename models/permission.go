package models

type Level string

const (
	Supervisor Level = "Supervisor"
	Control    Level = "Control"
	Monitor    Level = "Monitor"
)

type Permission struct {
	PermissionID string `json:"permission_id"`
	Level        Level  `json:"level"`
}
