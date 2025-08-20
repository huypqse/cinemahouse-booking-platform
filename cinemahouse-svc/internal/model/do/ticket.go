// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Ticket is the golang structure of table ticket for DAO operations like Where/Data.
type Ticket struct {
	g.Meta             `orm:"table:ticket, do:true"`
	Id                 interface{} //
	Status             interface{} //
	ScreeningSessionId interface{} //
	TicketPriceId      interface{} //
	BillId             interface{} //
	SeatId             interface{} //
}
