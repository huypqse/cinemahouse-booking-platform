package booking

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/booking/booking"
)

func (c *ControllerBooking) BookingGetTicket(ctx context.Context, req *booking.BookingGetTicketReq) (res *booking.BookingGetTicketRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
