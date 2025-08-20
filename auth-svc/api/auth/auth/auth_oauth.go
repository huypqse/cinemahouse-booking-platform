package auth

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OAuth authentication APIs
type AuthGoogleAuthReq struct {
	g.Meta      `path:"/auth/google" tags:"Auth" method:"post" summary:"Google OAuth authentication"`
	AccessToken string `json:"access_token" v:"required#Google access token is required"`
	RedirectUri string `json:"redirect_uri"`
}

type AuthGoogleAuthRes struct {
	UserId       int64  `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	IsNewUser    bool   `json:"is_new_user"`
}

type AuthFacebookAuthReq struct {
	g.Meta      `path:"/auth/facebook" tags:"Auth" method:"post" summary:"Facebook OAuth authentication"`
	AccessToken string `json:"access_token" v:"required#Facebook access token is required"`
	RedirectUri string `json:"redirect_uri"`
}

type AuthFacebookAuthRes struct {
	UserId       int64  `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	IsNewUser    bool   `json:"is_new_user"`
}
