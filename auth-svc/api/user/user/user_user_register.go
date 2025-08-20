package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserRegisterReq struct {
	g.Meta      `path:"/user/register" tags:"User" method:"post" summary:"User registration"`
	Username    string `json:"username" v:"required|length:3,50#Username is required|Username must be between 3 and 50 characters"`
	Email       string `json:"email" v:"required|email#Email is required|Invalid email format"`
	Password    string `json:"password" v:"required|length:6,32#Password is required|Password must be between 6 and 32 characters"`
	PhoneNumber string `json:"phone_number" v:"required|phone#Phone number is required|Invalid phone number format"`
	Gender      string `json:"gender" v:"in:MALE,FEMALE,OTHER#Gender must be MALE, FEMALE, or OTHER"`
	Dob         string `json:"dob" v:"date#Invalid date format for date of birth"`
}

type UserRegisterRes struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}
