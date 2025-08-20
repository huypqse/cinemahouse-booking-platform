// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"auth-svc/api/admin/admin"
)

type IAdminV1 interface {
	// User Management
	GetUsers(ctx context.Context, req *admin.AdminGetUsersReq) (res *admin.AdminGetUsersRes, err error)
	GetUser(ctx context.Context, req *admin.AdminGetUserReq) (res *admin.AdminGetUserRes, err error)
	UpdateUser(ctx context.Context, req *admin.AdminUpdateUserReq) (res *admin.AdminUpdateUserRes, err error)
	DeleteUser(ctx context.Context, req *admin.AdminDeleteUserReq) (res *admin.AdminDeleteUserRes, err error)
	BanUser(ctx context.Context, req *admin.AdminBanUserReq) (res *admin.AdminBanUserRes, err error)
	UnbanUser(ctx context.Context, req *admin.AdminUnbanUserReq) (res *admin.AdminUnbanUserRes, err error)
	
	// Role Management
	GetRoles(ctx context.Context, req *admin.AdminGetRolesReq) (res *admin.AdminGetRolesRes, err error)
	CreateRole(ctx context.Context, req *admin.AdminCreateRoleReq) (res *admin.AdminCreateRoleRes, err error)
	UpdateRole(ctx context.Context, req *admin.AdminUpdateRoleReq) (res *admin.AdminUpdateRoleRes, err error)
	DeleteRole(ctx context.Context, req *admin.AdminDeleteRoleReq) (res *admin.AdminDeleteRoleRes, err error)
	AssignRole(ctx context.Context, req *admin.AdminAssignRoleReq) (res *admin.AdminAssignRoleRes, err error)
	RemoveRole(ctx context.Context, req *admin.AdminRemoveRoleReq) (res *admin.AdminRemoveRoleRes, err error)
	
	// Permission Management
	GetPermissions(ctx context.Context, req *admin.AdminGetPermissionsReq) (res *admin.AdminGetPermissionsRes, err error)
	CreatePermission(ctx context.Context, req *admin.AdminCreatePermissionReq) (res *admin.AdminCreatePermissionRes, err error)
	UpdatePermission(ctx context.Context, req *admin.AdminUpdatePermissionReq) (res *admin.AdminUpdatePermissionRes, err error)
	DeletePermission(ctx context.Context, req *admin.AdminDeletePermissionReq) (res *admin.AdminDeletePermissionRes, err error)
	AssignPermission(ctx context.Context, req *admin.AdminAssignPermissionReq) (res *admin.AdminAssignPermissionRes, err error)
	RemovePermission(ctx context.Context, req *admin.AdminRemovePermissionReq) (res *admin.AdminRemovePermissionRes, err error)
	
	// System Management
	GetSystemStats(ctx context.Context, req *admin.AdminGetSystemStatsReq) (res *admin.AdminGetSystemStatsRes, err error)
	GetInvalidatedTokens(ctx context.Context, req *admin.AdminGetInvalidatedTokensReq) (res *admin.AdminGetInvalidatedTokensRes, err error)
	CleanupExpiredTokens(ctx context.Context, req *admin.AdminCleanupExpiredTokensReq) (res *admin.AdminCleanupExpiredTokensRes, err error)
}
