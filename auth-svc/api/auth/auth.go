// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"auth-svc/api/auth/auth"
)

type IAuthV1 interface {
	// Token validation and middleware
	ValidateToken(ctx context.Context, req *auth.AuthValidateTokenReq) (res *auth.AuthValidateTokenRes, err error)
	CheckPermission(ctx context.Context, req *auth.AuthCheckPermissionReq) (res *auth.AuthCheckPermissionRes, err error)
	GetUserRoles(ctx context.Context, req *auth.AuthGetUserRolesReq) (res *auth.AuthGetUserRolesRes, err error)
	GetUserPermissions(ctx context.Context, req *auth.AuthGetUserPermissionsReq) (res *auth.AuthGetUserPermissionsRes, err error)
	
	// OAuth and external authentication
	GoogleAuth(ctx context.Context, req *auth.AuthGoogleAuthReq) (res *auth.AuthGoogleAuthRes, err error)
	FacebookAuth(ctx context.Context, req *auth.AuthFacebookAuthReq) (res *auth.AuthFacebookAuthRes, err error)
	
	// Two-factor authentication
	EnableTwoFactor(ctx context.Context, req *auth.AuthEnableTwoFactorReq) (res *auth.AuthEnableTwoFactorRes, err error)
	DisableTwoFactor(ctx context.Context, req *auth.AuthDisableTwoFactorReq) (res *auth.AuthDisableTwoFactorRes, err error)
	VerifyTwoFactor(ctx context.Context, req *auth.AuthVerifyTwoFactorReq) (res *auth.AuthVerifyTwoFactorRes, err error)
}
