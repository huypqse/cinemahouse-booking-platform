package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Permission Management APIs
type AdminGetPermissionsReq struct {
	g.Meta `path:"/admin/permissions" tags:"Admin" method:"get" summary:"Get all permissions"`
}

type AdminGetPermissionsRes struct {
	Permissions []PermissionInfo `json:"permissions"`
}

type PermissionInfo struct {
	PermissionId int64    `json:"permission_id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Roles        []string `json:"roles"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type AdminCreatePermissionReq struct {
	g.Meta      `path:"/admin/permissions" tags:"Admin" method:"post" summary:"Create new permission"`
	Name        string `json:"name" v:"required|length:3,100#Permission name is required|Permission name must be between 3 and 100 characters"`
	Description string `json:"description" v:"length:0,255#Description must be less than 255 characters"`
}

type AdminCreatePermissionRes struct {
	PermissionId int64  `json:"permission_id"`
	Message      string `json:"message"`
}

type AdminUpdatePermissionReq struct {
	g.Meta       `path:"/admin/permissions/{permission_id}" tags:"Admin" method:"put" summary:"Update permission"`
	PermissionId int64  `json:"permission_id" v:"required|min:1#Permission ID is required|Permission ID must be positive"`
	Name         string `json:"name" v:"length:3,100#Permission name must be between 3 and 100 characters"`
	Description  string `json:"description" v:"length:0,255#Description must be less than 255 characters"`
}

type AdminUpdatePermissionRes struct {
	Message string `json:"message"`
}

type AdminDeletePermissionReq struct {
	g.Meta       `path:"/admin/permissions/{permission_id}" tags:"Admin" method:"delete" summary:"Delete permission"`
	PermissionId int64 `json:"permission_id" v:"required|min:1#Permission ID is required|Permission ID must be positive"`
}

type AdminDeletePermissionRes struct {
	Message string `json:"message"`
}

type AdminAssignPermissionReq struct {
	g.Meta       `path:"/admin/roles/{role_id}/permissions" tags:"Admin" method:"post" summary:"Assign permission to role"`
	RoleId       int64 `json:"role_id" v:"required|min:1#Role ID is required|Role ID must be positive"`
	PermissionId int64 `json:"permission_id" v:"required|min:1#Permission ID is required|Permission ID must be positive"`
}

type AdminAssignPermissionRes struct {
	Message string `json:"message"`
}

type AdminRemovePermissionReq struct {
	g.Meta       `path:"/admin/roles/{role_id}/permissions/{permission_id}" tags:"Admin" method:"delete" summary:"Remove permission from role"`
	RoleId       int64 `json:"role_id" v:"required|min:1#Role ID is required|Role ID must be positive"`
	PermissionId int64 `json:"permission_id" v:"required|min:1#Permission ID is required|Permission ID must be positive"`
}

type AdminRemovePermissionRes struct {
	Message string `json:"message"`
}
