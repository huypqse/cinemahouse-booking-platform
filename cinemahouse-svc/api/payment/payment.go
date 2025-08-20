// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package payment

import (
	"context"

	"cinemahouse-svc/api/payment/payment"
)

type IPaymentPayment interface {
	PaymentCreatePayment(ctx context.Context, req *payment.PaymentCreatePaymentReq) (res *payment.PaymentCreatePaymentRes, err error)
	PaymentProcessPayment(ctx context.Context, req *payment.PaymentProcessPaymentReq) (res *payment.PaymentProcessPaymentRes, err error)
	PaymentGetPayment(ctx context.Context, req *payment.PaymentGetPaymentReq) (res *payment.PaymentGetPaymentRes, err error)
	PaymentGetPayments(ctx context.Context, req *payment.PaymentGetPaymentsReq) (res *payment.PaymentGetPaymentsRes, err error)
	PaymentRefundPayment(ctx context.Context, req *payment.PaymentRefundPaymentReq) (res *payment.PaymentRefundPaymentRes, err error)
	PaymentGetTicketPrices(ctx context.Context, req *payment.PaymentGetTicketPricesReq) (res *payment.PaymentGetTicketPricesRes, err error)
	PaymentCreateTicketPrice(ctx context.Context, req *payment.PaymentCreateTicketPriceReq) (res *payment.PaymentCreateTicketPriceRes, err error)
	PaymentUpdateTicketPrice(ctx context.Context, req *payment.PaymentUpdateTicketPriceReq) (res *payment.PaymentUpdateTicketPriceRes, err error)
	PaymentDeleteTicketPrice(ctx context.Context, req *payment.PaymentDeleteTicketPriceReq) (res *payment.PaymentDeleteTicketPriceRes, err error)
}
