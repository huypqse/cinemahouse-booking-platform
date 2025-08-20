// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"auth-svc/api/auth/auth"
)

type IAuthAuth interface {
	AuthGoogleAuth(ctx context.Context, req *auth.AuthGoogleAuthReq) (res *auth.AuthGoogleAuthRes, err error)
	AuthFacebookAuth(ctx context.Context, req *auth.AuthFacebookAuthReq) (res *auth.AuthFacebookAuthRes, err error)
	AuthValidateToken(ctx context.Context, req *auth.AuthValidateTokenReq) (res *auth.AuthValidateTokenRes, err error)
	AuthCheckPermission(ctx context.Context, req *auth.AuthCheckPermissionReq) (res *auth.AuthCheckPermissionRes, err error)
	AuthGetUserRoles(ctx context.Context, req *auth.AuthGetUserRolesReq) (res *auth.AuthGetUserRolesRes, err error)
	AuthGetUserPermissions(ctx context.Context, req *auth.AuthGetUserPermissionsReq) (res *auth.AuthGetUserPermissionsRes, err error)
	AuthEnableTwoFactor(ctx context.Context, req *auth.AuthEnableTwoFactorReq) (res *auth.AuthEnableTwoFactorRes, err error)
	AuthDisableTwoFactor(ctx context.Context, req *auth.AuthDisableTwoFactorReq) (res *auth.AuthDisableTwoFactorRes, err error)
	AuthVerifyTwoFactor(ctx context.Context, req *auth.AuthVerifyTwoFactorReq) (res *auth.AuthVerifyTwoFactorRes, err error)
}
