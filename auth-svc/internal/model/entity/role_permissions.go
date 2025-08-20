// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RolePermissions is the golang structure for table role_permissions.
type RolePermissions struct {
	RoleId       int64 `json:"roleId"       orm:"role_id"       description:""`
	PermissionId int64 `json:"permissionId" orm:"permission_id" description:""`
}
