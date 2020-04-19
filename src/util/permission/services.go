package permission

import (
	"log"

	"fmt"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/modulegroup"
	"github.com/KitaPDev/fogfarms-server/src/util/permission/repository"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
)

func GetUserModuleGroupPermissions(userIDs []int, moduleGroupIDs []int) (map[string]map[string]int, error) {
	if len(userIDs) == 0 || len(moduleGroupIDs) == 0 {
		return make(map[string]map[string]int), nil
	}

	userModuleGroupPermissions := make(map[string]map[string]int)

	permissions, err := repository.GetAllPermissions()
	if err != nil {
		return make(map[string]map[string]int), err
	}
	log.Println("Variable permissions in GetuserModuleGroupPermissions", permissions)

	users, err := user.GetUsersByID(userIDs)
	if err != nil {
		return make(map[string]map[string]int), err
	}
	log.Println("Variable users in GetuserModuleGroupPermissions", users)

	moduleGroups, err := modulegroup.GetModuleGroupsByIDs(moduleGroupIDs)
	if err != nil {
		return make(map[string]map[string]int), err
	}
	log.Println("Variable moduleGroups in GetuserModuleGroupPermissions", moduleGroups)

	fmt.Println("permissions:\n", permissions)
	fmt.Println("users:\n", users)
	fmt.Println("moduleGroups:\n", moduleGroups)

	fGetUsername := func(userID int) string {
		for _, uTemp := range users {
			if uTemp.UserID == userID {
				return uTemp.Username
			}
		}
		return ""
	}

	fGetModuleGroupLabel := func(moduleGroupID int) string {
		for _, mg := range moduleGroups {
			if mg.ModuleGroupID == moduleGroupID {
				return mg.ModuleGroupLabel
			}
		}
		return ""
	}

	fGetPermission := func(userID int, moduleGroupID int) int {
		for _, p := range permissions {
			if p.UserID == userID && p.ModuleGroupID == moduleGroupID {
				return p.PermissionLevel
			}
		}
		return 0
	}
	log.Println("Variable moduleGroups in GetuserModuleGroupPermissions", moduleGroups)

	for _, uid := range userIDs {
		userModuleGroupPermissions[fGetUsername(uid)] = make(map[string]int)

		for mgid := range moduleGroupIDs {
			userModuleGroupPermissions[fGetUsername(uid)][fGetModuleGroupLabel(mgid)] =
				fGetPermission(uid, mgid)
		}

	}
	fmt.Println("userModuleGroupPermissions:\n", userModuleGroupPermissions)

	return userModuleGroupPermissions, nil
}

func AssignUserModuleGroupPermission(username string, ModuleGroupLabel string, permissionLevel int) error {
	return repository.AssignUserModuleGroupPermission(username, ModuleGroupLabel, permissionLevel)
}

func GetSupervisorModuleGroups(user *models.User) ([]models.ModuleGroup, error) {
	mapModuleGroupPermissionLevels, err := repository.GetAssignedModuleGroupsWithPermissionLevel(user.UserID, 3)
	if err != nil {
		return nil, err
	}

	moduleGroups := make([]models.ModuleGroup, 0, len(mapModuleGroupPermissionLevels))
	for k := range mapModuleGroupPermissionLevels {
		moduleGroups = append(moduleGroups, k)
	}

	return moduleGroups, nil
}

func GetAssignedModuleGroups(user *models.User) (map[models.ModuleGroup]int, error) {
	return repository.GetAssignedModuleGroupsWithPermissionLevel(user.UserID, -1)
}

func PopulateUserManagementPage(u *models.User) (map[string]map[string]int, error) {
	return repository.PopulateUserManagementPage(u)
}
