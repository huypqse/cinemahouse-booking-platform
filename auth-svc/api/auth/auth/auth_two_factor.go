package auth

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Two-factor authentication APIs
type AuthEnableTwoFactorReq struct {
	g.Meta   `path:"/auth/2fa/enable" tags:"Auth" method:"post" summary:"Enable two-factor authentication"`
	Password string `json:"password" v:"required#Password is required"`
}

type AuthEnableTwoFactorRes struct {
	SecretKey string `json:"secret_key"`
	QrCode    string `json:"qr_code"`
	Message   string `json:"message"`
}

type AuthDisableTwoFactorReq struct {
	g.Meta   `path:"/auth/2fa/disable" tags:"Auth" method:"post" summary:"Disable two-factor authentication"`
	Password string `json:"password" v:"required#Password is required"`
	Code     string `json:"code" v:"required|length:6,6#2FA code is required|2FA code must be exactly 6 digits"`
}

type AuthDisableTwoFactorRes struct {
	Message string `json:"message"`
}

type AuthVerifyTwoFactorReq struct {
	g.Meta `path:"/auth/2fa/verify" tags:"Auth" method:"post" summary:"Verify two-factor authentication code"`
	Code   string `json:"code" v:"required|length:6,6#2FA code is required|2FA code must be exactly 6 digits"`
}

type AuthVerifyTwoFactorRes struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}
