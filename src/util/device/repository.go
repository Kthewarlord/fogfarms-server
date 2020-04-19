package device

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetModuleGroupDevices(moduleGroupID int) ([]models.Device, error)
	ToggleDevice(deviceID int) error
}
