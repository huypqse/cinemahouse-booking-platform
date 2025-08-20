// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Seat is the golang structure of table seat for DAO operations like Where/Data.
type Seat struct {
	g.Meta          `orm:"table:seat, do:true"`
	Id              interface{} //
	SeatRow         interface{} //
	SeatColumn      interface{} //
	Status          interface{} //
	ScreeningRoomId interface{} //
	SeatTypeId      interface{} //
}
