// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// InvalidatedTokens is the golang structure for table invalidated_tokens.
type InvalidatedTokens struct {
	Id         string      `json:"id"         orm:"id"          description:""`
	ExpiryTime *gtime.Time `json:"expiryTime" orm:"expiry_time" description:""`
}
