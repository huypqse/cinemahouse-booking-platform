package payment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/payment/payment"
)

func (c *ControllerPayment) PaymentUpdateTicketPrice(ctx context.Context, req *payment.PaymentUpdateTicketPriceReq) (res *payment.PaymentUpdateTicketPriceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
