package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Email verification APIs
type UserVerifyEmailReq struct {
	g.Meta `path:"/user/verify-email" tags:"User" method:"post" summary:"Verify email with OTP"`
	Email  string `json:"email" v:"required|email#Email is required|Invalid email format"`
	Otp    string `json:"otp" v:"required|length:6,6#OTP is required|OTP must be exactly 6 digits"`
}

type UserVerifyEmailRes struct {
	Message string `json:"message"`
}

type UserResendVerificationReq struct {
	g.Meta `path:"/user/resend-verification" tags:"User" method:"post" summary:"Resend email verification OTP"`
	Email  string `json:"email" v:"required|email#Email is required|Invalid email format"`
}

type UserResendVerificationRes struct {
	Message string `json:"message"`
}

// Refresh token API
type UserRefreshTokenReq struct {
	g.Meta       `path:"/user/refresh-token" tags:"User" method:"post" summary:"Refresh access token"`
	RefreshToken string `json:"refresh_token" v:"required#Refresh token is required"`
}

type UserRefreshTokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

// Account status management
type UserDeactivateAccountReq struct {
	g.Meta   `path:"/user/deactivate" tags:"User" method:"post" summary:"Deactivate user account"`
	Password string `json:"password" v:"required#Password is required to deactivate account"`
}

type UserDeactivateAccountRes struct {
	Message string `json:"message"`
}

type UserReactivateAccountReq struct {
	g.Meta `path:"/user/reactivate" tags:"User" method:"post" summary:"Reactivate user account"`
	Email  string `json:"email" v:"required|email#Email is required|Invalid email format"`
	Otp    string `json:"otp" v:"required|length:6,6#OTP is required|OTP must be exactly 6 digits"`
}

type UserReactivateAccountRes struct {
	Message string `json:"message"`
}
