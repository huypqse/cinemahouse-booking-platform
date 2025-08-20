package booking

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/booking/booking"
)

func (c *ControllerBooking) BookingGetTickets(ctx context.Context, req *booking.BookingGetTicketsReq) (res *booking.BookingGetTicketsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
