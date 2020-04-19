package module

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/models/outputs"
	"github.com/KitaPDev/fogfarms-server/src/util/module/repository"
)

func GetModulesByModuleGroupIDs(moduleGroupIDs []int) ([]models.Module, error) {
	return repository.GetModulesByModuleGroupIDs(moduleGroupIDs)
}

func GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs []int) ([]outputs.ModuleOutput, error) {
	return repository.GetModulesByModuleGroupIDsForModuleManagement(moduleGroupIDs)
}

func AssignModulesToModuleGroup(moduleGroupID int, moduleIDs []int) error {
	return repository.AssignModulesToModuleGroup(moduleGroupID, moduleIDs)
}