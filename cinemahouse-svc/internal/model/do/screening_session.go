// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ScreeningSession is the golang structure of table screening_session for DAO operations like Where/Data.
type ScreeningSession struct {
	g.Meta                 `orm:"table:screening_session, do:true"`
	Id                     interface{} //
	StartDate              *gtime.Time //
	StartTime              *gtime.Time //
	EndTime                *gtime.Time //
	ScreeningSessionStatus interface{} //
	MovieId                interface{} //
	ScreeningRoomId        interface{} //
}
