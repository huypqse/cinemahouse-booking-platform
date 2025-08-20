// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:user, do:true"`
	Id            interface{} //
	Username      interface{} //
	Gender        interface{} //
	PhoneNumber   interface{} //
	Dob           *gtime.Time //
	Otp           interface{} //
	OtpExpiryDate *gtime.Time //
	Email         interface{} //
	Password      interface{} //
	Account       interface{} //
	CommentId     interface{} //
}
