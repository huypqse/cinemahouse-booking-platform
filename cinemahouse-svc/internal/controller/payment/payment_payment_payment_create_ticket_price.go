package payment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/payment/payment"
)

func (c *ControllerPayment) PaymentCreateTicketPrice(ctx context.Context, req *payment.PaymentCreateTicketPriceReq) (res *payment.PaymentCreateTicketPriceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
