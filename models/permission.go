package models

type Permission struct {
	UserID          int `json:"user_id"`
	ModuleGroupID   int `json:"module_group_id"`
	PermissionLevel int `json:"permission_level"`
}
