// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"auth-svc/api/admin/admin"
)

type IAdminAdmin interface {
	AdminGetPermissions(ctx context.Context, req *admin.AdminGetPermissionsReq) (res *admin.AdminGetPermissionsRes, err error)
	AdminCreatePermission(ctx context.Context, req *admin.AdminCreatePermissionReq) (res *admin.AdminCreatePermissionRes, err error)
	AdminUpdatePermission(ctx context.Context, req *admin.AdminUpdatePermissionReq) (res *admin.AdminUpdatePermissionRes, err error)
	AdminDeletePermission(ctx context.Context, req *admin.AdminDeletePermissionReq) (res *admin.AdminDeletePermissionRes, err error)
	AdminAssignPermission(ctx context.Context, req *admin.AdminAssignPermissionReq) (res *admin.AdminAssignPermissionRes, err error)
	AdminRemovePermission(ctx context.Context, req *admin.AdminRemovePermissionReq) (res *admin.AdminRemovePermissionRes, err error)
	AdminGetRoles(ctx context.Context, req *admin.AdminGetRolesReq) (res *admin.AdminGetRolesRes, err error)
	AdminCreateRole(ctx context.Context, req *admin.AdminCreateRoleReq) (res *admin.AdminCreateRoleRes, err error)
	AdminUpdateRole(ctx context.Context, req *admin.AdminUpdateRoleReq) (res *admin.AdminUpdateRoleRes, err error)
	AdminDeleteRole(ctx context.Context, req *admin.AdminDeleteRoleReq) (res *admin.AdminDeleteRoleRes, err error)
	AdminAssignRole(ctx context.Context, req *admin.AdminAssignRoleReq) (res *admin.AdminAssignRoleRes, err error)
	AdminRemoveRole(ctx context.Context, req *admin.AdminRemoveRoleReq) (res *admin.AdminRemoveRoleRes, err error)
	AdminGetSystemStats(ctx context.Context, req *admin.AdminGetSystemStatsReq) (res *admin.AdminGetSystemStatsRes, err error)
	AdminGetInvalidatedTokens(ctx context.Context, req *admin.AdminGetInvalidatedTokensReq) (res *admin.AdminGetInvalidatedTokensRes, err error)
	AdminCleanupExpiredTokens(ctx context.Context, req *admin.AdminCleanupExpiredTokensReq) (res *admin.AdminCleanupExpiredTokensRes, err error)
	AdminGetAuditLogs(ctx context.Context, req *admin.AdminGetAuditLogsReq) (res *admin.AdminGetAuditLogsRes, err error)
	AdminGetUsers(ctx context.Context, req *admin.AdminGetUsersReq) (res *admin.AdminGetUsersRes, err error)
	AdminGetUser(ctx context.Context, req *admin.AdminGetUserReq) (res *admin.AdminGetUserRes, err error)
	AdminUpdateUser(ctx context.Context, req *admin.AdminUpdateUserReq) (res *admin.AdminUpdateUserRes, err error)
	AdminDeleteUser(ctx context.Context, req *admin.AdminDeleteUserReq) (res *admin.AdminDeleteUserRes, err error)
	AdminBanUser(ctx context.Context, req *admin.AdminBanUserReq) (res *admin.AdminBanUserRes, err error)
	AdminUnbanUser(ctx context.Context, req *admin.AdminUnbanUserReq) (res *admin.AdminUnbanUserRes, err error)
}
