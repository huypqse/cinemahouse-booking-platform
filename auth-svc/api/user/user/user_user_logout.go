package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserLogoutReq struct {
	g.Meta `path:"/user/logout" tags:"User" method:"post" summary:"User logout"`
}

type UserLogoutRes struct {
	Message string `json:"message"`
}
