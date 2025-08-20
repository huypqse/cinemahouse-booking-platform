// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ScreeningRoom is the golang structure of table screening_room for DAO operations like Where/Data.
type ScreeningRoom struct {
	g.Meta              `orm:"table:screening_room, do:true"`
	Id                  interface{} //
	Name                interface{} //
	ScreeningRoomStatus interface{} //
	RoomtypeId          interface{} //
	CinemaId            interface{} //
}
