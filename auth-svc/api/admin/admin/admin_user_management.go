package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User Management APIs
type AdminGetUsersReq struct {
	g.Meta   `path:"/admin/users" tags:"Admin" method:"get" summary:"Get all users with pagination"`
	Page     int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	Search   string `json:"search"`
	Status   string `json:"status" v:"in:ACTIVE,INACTIVE,BANNED#Status must be ACTIVE, INACTIVE, or BANNED"`
}

type AdminGetUsersRes struct {
	Users      []UserInfo `json:"users"`
	Total      int64      `json:"total"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}

type UserInfo struct {
	UserId      int64    `json:"user_id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	Gender      string   `json:"gender"`
	Dob         string   `json:"dob"`
	Account     string   `json:"account"`
	Roles       []string `json:"roles"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type AdminGetUserReq struct {
	g.Meta `path:"/admin/users/{user_id}" tags:"Admin" method:"get" summary:"Get user by ID"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
}

type AdminGetUserRes struct {
	User UserInfo `json:"user"`
}

type AdminUpdateUserReq struct {
	g.Meta      `path:"/admin/users/{user_id}" tags:"Admin" method:"put" summary:"Update user information"`
	UserId      int64  `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
	Username    string `json:"username" v:"length:3,50#Username must be between 3 and 50 characters"`
	Email       string `json:"email" v:"email#Invalid email format"`
	PhoneNumber string `json:"phone_number" v:"phone#Invalid phone number format"`
	Gender      string `json:"gender" v:"in:MALE,FEMALE,OTHER#Gender must be MALE, FEMALE, or OTHER"`
	Dob         string `json:"dob" v:"date#Invalid date format"`
	Account     string `json:"account" v:"in:ACTIVE,INACTIVE,BANNED#Account status must be ACTIVE, INACTIVE, or BANNED"`
}

type AdminUpdateUserRes struct {
	Message string `json:"message"`
}

type AdminDeleteUserReq struct {
	g.Meta `path:"/admin/users/{user_id}" tags:"Admin" method:"delete" summary:"Delete user"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
}

type AdminDeleteUserRes struct {
	Message string `json:"message"`
}

type AdminBanUserReq struct {
	g.Meta `path:"/admin/users/{user_id}/ban" tags:"Admin" method:"post" summary:"Ban user"`
	UserId int64  `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
	Reason string `json:"reason" v:"required#Ban reason is required"`
}

type AdminBanUserRes struct {
	Message string `json:"message"`
}

type AdminUnbanUserReq struct {
	g.Meta `path:"/admin/users/{user_id}/unban" tags:"Admin" method:"post" summary:"Unban user"`
	UserId int64 `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
}

type AdminUnbanUserRes struct {
	Message string `json:"message"`
}
