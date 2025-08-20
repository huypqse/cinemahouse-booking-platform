package payment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/payment/payment"
)

func (c *ControllerPayment) PaymentCreatePayment(ctx context.Context, req *payment.PaymentCreatePaymentReq) (res *payment.PaymentCreatePaymentRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
