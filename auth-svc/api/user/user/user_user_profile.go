package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" tags:"User" method:"get" summary:"Get user profile"`
}

type UserProfileRes struct {
	UserId      int64  `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Dob         string `json:"dob"`
	Account     string `json:"account"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserUpdateProfileReq struct {
	g.Meta      `path:"/user/profile" tags:"User" method:"put" summary:"Update user profile"`
	Username    string `json:"username" v:"length:3,50#Username must be between 3 and 50 characters"`
	PhoneNumber string `json:"phone_number" v:"phone#Invalid phone number format"`
	Gender      string `json:"gender" v:"in:MALE,FEMALE,OTHER#Gender must be MALE, FEMALE, or OTHER"`
	Dob         string `json:"dob" v:"date#Invalid date format for date of birth"`
}

type UserUpdateProfileRes struct {
	Message string `json:"message"`
}
