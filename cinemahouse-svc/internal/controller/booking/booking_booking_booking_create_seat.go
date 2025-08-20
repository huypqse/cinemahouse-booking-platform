package booking

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/booking/booking"
)

func (c *ControllerBooking) BookingCreateSeat(ctx context.Context, req *booking.BookingCreateSeatReq) (res *booking.BookingCreateSeatRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
