package module

import (
	"github.com/KitaPDev/fogfarms-server/models"
)

type Repository interface {
	CreateModule(moduleLabel string) error
	GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error)
	AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error
	GetModuleIDByToken(token string) (int, error)
	GetDeviceStatus(moduleID int) ([]bool, []bool, []bool, []bool, error)
	UpdateDeviceStatus(moduleID int, mixer []bool, solenoidValves []bool, led []bool,
		fogger []bool) error
}

