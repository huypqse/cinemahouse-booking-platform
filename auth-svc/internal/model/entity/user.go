// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id            int64       `json:"id"            orm:"id"              description:""`
	Username      string      `json:"username"      orm:"username"        description:""`
	Gender        string      `json:"gender"        orm:"gender"          description:""`
	PhoneNumber   string      `json:"phoneNumber"   orm:"phone_number"    description:""`
	Dob           *gtime.Time `json:"dob"           orm:"dob"             description:""`
	Otp           string      `json:"otp"           orm:"otp"             description:""`
	OtpExpiryDate *gtime.Time `json:"otpExpiryDate" orm:"otp_expiry_date" description:""`
	Email         string      `json:"email"         orm:"email"           description:""`
	Password      string      `json:"password"      orm:"password"        description:""`
	Account       string      `json:"account"       orm:"account"         description:""`
	CommentId     int64       `json:"commentId"     orm:"comment_id"      description:""`
}
