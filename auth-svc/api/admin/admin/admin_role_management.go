package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Role Management APIs
type AdminGetRolesReq struct {
	g.Meta `path:"/admin/roles" tags:"Admin" method:"get" summary:"Get all roles"`
}

type AdminGetRolesRes struct {
	Roles []RoleInfo `json:"roles"`
}

type RoleInfo struct {
	RoleId      int64    `json:"role_id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	UserCount   int64    `json:"user_count"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type AdminCreateRoleReq struct {
	g.Meta      `path:"/admin/roles" tags:"Admin" method:"post" summary:"Create new role"`
	Name        string  `json:"name" v:"required|length:3,100#Role name is required|Role name must be between 3 and 100 characters"`
	Permissions []int64 `json:"permissions"`
}

type AdminCreateRoleRes struct {
	RoleId  int64  `json:"role_id"`
	Message string `json:"message"`
}

type AdminUpdateRoleReq struct {
	g.Meta      `path:"/admin/roles/{role_id}" tags:"Admin" method:"put" summary:"Update role"`
	RoleId      int64   `json:"role_id" v:"required|min:1#Role ID is required|Role ID must be positive"`
	Name        string  `json:"name" v:"length:3,100#Role name must be between 3 and 100 characters"`
	Permissions []int64 `json:"permissions"`
}

type AdminUpdateRoleRes struct {
	Message string `json:"message"`
}

type AdminDeleteRoleReq struct {
	g.Meta `path:"/admin/roles/{role_id}" tags:"Admin" method:"delete" summary:"Delete role"`
	RoleId int64 `json:"role_id" v:"required|min:1#Role ID is required|Role ID must be positive"`
}

type AdminDeleteRoleRes struct {
	Message string `json:"message"`
}

type AdminAssignRoleReq struct {
	g.Meta `path:"/admin/users/{user_id}/roles" tags:"Admin" method:"post" summary:"Assign role to user"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
	RoleId int64 `json:"role_id" v:"required|min:1#Role ID is required|Role ID must be positive"`
}

type AdminAssignRoleRes struct {
	Message string `json:"message"`
}

type AdminRemoveRoleReq struct {
	g.Meta `path:"/admin/users/{user_id}/roles/{role_id}" tags:"Admin" method:"delete" summary:"Remove role from user"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
	RoleId int64 `json:"role_id" v:"required|min:1#Role ID is required|Role ID must be positive"`
}

type AdminRemoveRoleRes struct {
	Message string `json:"message"`
}
