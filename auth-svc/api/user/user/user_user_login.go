package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"User" method:"post" summary:"User login"`
	Email    string `json:"email" v:"required|email#Email is required|Invalid email format"`
	Password string `json:"password" v:"required#Password is required"`
}

type UserLoginRes struct {
	UserId      int64  `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
