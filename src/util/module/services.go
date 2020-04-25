package module

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/models/outputs"
	"github.com/KitaPDev/fogfarms-server/src/util/module/repository"
)

func CreateModule(moduleLabel string) error {
	return repository.CreateModule(moduleLabel)
}

func DeleteModule(moduleLabel string) error {
	return repository.DeleteModule(moduleLabel)
}

func EditModuleLabel(moduleID int, moduleLabel string) error {
	return repository.EditModuleLabel(moduleID, moduleLabel)
}
func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	return repository.GetModulesByModuleGroupIDs(moduleGroupIDs)
}

func GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs []int) ([]outputs.ModuleOutput, error) {
	return repository.GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs)
}

func AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error {
	return repository.AssignModulesToModuleGroup(moduleGroupID, moduleIDs)
}

func GetModuleIDByToken(token string) (int, error) {
	return repository.GetModuleIDByToken(token)
}

func UpdateDeviceStatus(moduleID int, mixer []bool, solenoidValves []bool, led []bool,
	fogger []bool) error {

	return repository.UpdateDeviceStatus(moduleID, mixer, solenoidValves, led, fogger)
}

func GetDeviceStatus(moduleID int) ([]bool, []bool, []bool, []bool, error) {
	return repository.GetDeviceStatus(moduleID)
}
