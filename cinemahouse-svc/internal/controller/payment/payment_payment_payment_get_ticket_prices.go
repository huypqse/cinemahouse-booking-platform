package payment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/payment/payment"
)

func (c *ControllerPayment) PaymentGetTicketPrices(ctx context.Context, req *payment.PaymentGetTicketPricesReq) (res *payment.PaymentGetTicketPricesRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
