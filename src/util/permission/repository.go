package permission

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllPermissions () ([]models.Permission, error)
	AssignUserToModuleGroup(userID int, moduleGroupID int, permissionLevel int) error
	GetAssignedModuleGroupsWithPermissionLevel(userID int, permissionLevel int) (map[*models.ModuleGroup]int, error)}