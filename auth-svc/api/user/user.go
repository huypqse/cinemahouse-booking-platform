// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"auth-svc/api/user/user"
)

type IUserV1 interface {
	Register(ctx context.Context, req *user.UserRegisterReq) (res *user.UserRegisterRes, err error)
	Login(ctx context.Context, req *user.UserLoginReq) (res *user.UserLoginRes, err error)
	Logout(ctx context.Context, req *user.UserLogoutReq) (res *user.UserLogoutRes, err error)
	Profile(ctx context.Context, req *user.UserProfileReq) (res *user.UserProfileRes, err error)
	UpdateProfile(ctx context.Context, req *user.UserUpdateProfileReq) (res *user.UserUpdateProfileRes, err error)
}
