package booking

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/booking/booking"
)

func (c *ControllerBooking) BookingGetBill(ctx context.Context, req *booking.BookingGetBillReq) (res *booking.BookingGetBillRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
