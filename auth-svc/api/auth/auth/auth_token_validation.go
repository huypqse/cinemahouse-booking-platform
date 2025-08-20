package auth

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Token validation APIs
type AuthValidateTokenReq struct {
	g.Meta `path:"/auth/validate-token" tags:"Auth" method:"post" summary:"Validate access token"`
	Token  string `json:"token" v:"required#Token is required"`
}

type AuthValidateTokenRes struct {
	Valid     bool   `json:"valid"`
	UserId    int64  `json:"user_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	ExpiresAt string `json:"expires_at,omitempty"`
	Message   string `json:"message,omitempty"`
}

type AuthCheckPermissionReq struct {
	g.Meta     `path:"/auth/check-permission" tags:"Auth" method:"post" summary:"Check user permission"`
	UserId     int64  `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
	Permission string `json:"permission" v:"required#Permission is required"`
}

type AuthCheckPermissionRes struct {
	HasPermission bool   `json:"has_permission"`
	Message       string `json:"message,omitempty"`
}

type AuthGetUserRolesReq struct {
	g.Meta `path:"/auth/user/{user_id}/roles" tags:"Auth" method:"get" summary:"Get user roles"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
}

type AuthGetUserRolesRes struct {
	UserId int64    `json:"user_id"`
	Roles  []string `json:"roles"`
}

type AuthGetUserPermissionsReq struct {
	g.Meta `path:"/auth/user/{user_id}/permissions" tags:"Auth" method:"get" summary:"Get user permissions"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
}

type AuthGetUserPermissionsRes struct {
	UserId      int64    `json:"user_id"`
	Permissions []string `json:"permissions"`
}
