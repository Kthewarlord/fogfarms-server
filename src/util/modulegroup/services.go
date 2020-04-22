package modulegroup

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup/repository"
)

func GetAllModuleGroups() ([]models.ModuleGroup, error) {
	moduleGroups, err := repository.GetAllModuleGroups()
	return moduleGroups, err
}

func GetModuleGroupsByIDs(moduleGroupIDs []int) ([]models.ModuleGroup, error) {
	moduleGroups, err := repository.GetModuleGroupsByIDs(moduleGroupIDs)
	return moduleGroups, err
}

func CreateModuleGroup(label string, plantID int, locationID int, tds float64, ph float64,
	humidity float64, lightsOn float64, lightsOff float64, onAuto bool) error {

	return repository.CreateModuleGroup(label, plantID, locationID, tds, ph, humidity,
		lightsOn, lightsOff, onAuto)
}

func DeleteModuleGroup(moduleGroupID int) error {
	return repository.DeleteModuleGroup(moduleGroupID)
}

func EditModuleGroupLabel(moduleGroupID int, moduleGroupLabel string) error {
	return repository.EditModuleGroupLabel(moduleGroupID, moduleGroupLabel)
}

func ChangePlant(moduleGroupID int, plantID int) error {
	return repository.ChangePlant(moduleGroupID, plantID)
}

func ToggleAuto(moduleGroupID int) error {
	return repository.ToggleAuto(moduleGroupID)
}

func SetEnvironmentParameters(moduleGroupID int, humidity float64, ph float64, tds float64,
	lightsOnHour float64, lightsOffHour float64) error {

	return repository.SetEnvironmentParameters(moduleGroupID, humidity, ph, tds, lightsOnHour,
		lightsOffHour)
}

func ResetTimer(moduleGroupID int) error {
	return repository.ResetTimer(moduleGroupID)
}

func GetOnAutoByModuleID(moduleID int) (bool, error) {
	return repository.GetOnAutoByModuleID(moduleID)
}
func GetModuleGroupsByLabelMatch(moduleGroupLabel string)([]models.ModuleGroup,error){
	return repository.GetModuleGroupsByLabelMatch(moduleGroupLabel)
}
