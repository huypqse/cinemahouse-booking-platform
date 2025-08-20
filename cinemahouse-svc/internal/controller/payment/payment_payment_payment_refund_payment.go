package payment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/payment/payment"
)

func (c *ControllerPayment) PaymentRefundPayment(ctx context.Context, req *payment.PaymentRefundPaymentReq) (res *payment.PaymentRefundPaymentRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
