package payment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/payment/payment"
)

func (c *ControllerPayment) PaymentGetPayments(ctx context.Context, req *payment.PaymentGetPaymentsReq) (res *payment.PaymentGetPaymentsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
