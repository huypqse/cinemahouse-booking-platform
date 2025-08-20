// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"auth-svc/api/user/user"
)

type IUserUser interface {
	UserVerifyEmail(ctx context.Context, req *user.UserVerifyEmailReq) (res *user.UserVerifyEmailRes, err error)
	UserResendVerification(ctx context.Context, req *user.UserResendVerificationReq) (res *user.UserResendVerificationRes, err error)
	UserRefreshToken(ctx context.Context, req *user.UserRefreshTokenReq) (res *user.UserRefreshTokenRes, err error)
	UserDeactivateAccount(ctx context.Context, req *user.UserDeactivateAccountReq) (res *user.UserDeactivateAccountRes, err error)
	UserReactivateAccount(ctx context.Context, req *user.UserReactivateAccountReq) (res *user.UserReactivateAccountRes, err error)
	UserLogin(ctx context.Context, req *user.UserLoginReq) (res *user.UserLoginRes, err error)
	UserLogout(ctx context.Context, req *user.UserLogoutReq) (res *user.UserLogoutRes, err error)
	UserProfile(ctx context.Context, req *user.UserProfileReq) (res *user.UserProfileRes, err error)
	UserUpdateProfile(ctx context.Context, req *user.UserUpdateProfileReq) (res *user.UserUpdateProfileRes, err error)
	UserRegister(ctx context.Context, req *user.UserRegisterReq) (res *user.UserRegisterRes, err error)
}
