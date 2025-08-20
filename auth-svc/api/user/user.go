// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"auth-svc/api/user/user"
)

type IUserV1 interface {
	// Basic authentication
	Register(ctx context.Context, req *user.UserRegisterReq) (res *user.UserRegisterRes, err error)
	Login(ctx context.Context, req *user.UserLoginReq) (res *user.UserLoginRes, err error)
	Logout(ctx context.Context, req *user.UserLogoutReq) (res *user.UserLogoutRes, err error)
	
	// Profile management
	Profile(ctx context.Context, req *user.UserProfileReq) (res *user.UserProfileRes, err error)
	UpdateProfile(ctx context.Context, req *user.UserUpdateProfileReq) (res *user.UserUpdateProfileRes, err error)
	
	// Email verification
	VerifyEmail(ctx context.Context, req *user.UserVerifyEmailReq) (res *user.UserVerifyEmailRes, err error)
	ResendVerification(ctx context.Context, req *user.UserResendVerificationReq) (res *user.UserResendVerificationRes, err error)
	
	// Token management
	RefreshToken(ctx context.Context, req *user.UserRefreshTokenReq) (res *user.UserRefreshTokenRes, err error)
	
	// Account management
	DeactivateAccount(ctx context.Context, req *user.UserDeactivateAccountReq) (res *user.UserDeactivateAccountRes, err error)
	ReactivateAccount(ctx context.Context, req *user.UserReactivateAccountReq) (res *user.UserReactivateAccountRes, err error)
}
