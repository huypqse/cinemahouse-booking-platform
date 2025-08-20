// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// InvalidatedTokens is the golang structure of table invalidated_tokens for DAO operations like Where/Data.
type InvalidatedTokens struct {
	g.Meta     `orm:"table:invalidated_tokens, do:true"`
	Id         interface{} //
	ExpiryTime *gtime.Time //
}
