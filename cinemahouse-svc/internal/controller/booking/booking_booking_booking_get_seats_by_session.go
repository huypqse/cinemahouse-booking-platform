package booking

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/booking/booking"
)

func (c *ControllerBooking) BookingGetSeatsBySession(ctx context.Context, req *booking.BookingGetSeatsBySessionReq) (res *booking.BookingGetSeatsBySessionRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
