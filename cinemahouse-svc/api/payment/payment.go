// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package payment

import (
	"context"

	"cinemahouse-svc/api/payment/payment"
)

type IPaymentV1 interface {
	// Payment processing
	CreatePayment(ctx context.Context, req *payment.PaymentCreatePaymentReq) (res *payment.PaymentCreatePaymentRes, err error)
	ProcessPayment(ctx context.Context, req *payment.PaymentProcessPaymentReq) (res *payment.PaymentProcessPaymentRes, err error)
	GetPayment(ctx context.Context, req *payment.PaymentGetPaymentReq) (res *payment.PaymentGetPaymentRes, err error)
	GetPayments(ctx context.Context, req *payment.PaymentGetPaymentsReq) (res *payment.PaymentGetPaymentsRes, err error)
	RefundPayment(ctx context.Context, req *payment.PaymentRefundPaymentReq) (res *payment.PaymentRefundPaymentRes, err error)
	
	// Ticket price management
	GetTicketPrices(ctx context.Context, req *payment.PaymentGetTicketPricesReq) (res *payment.PaymentGetTicketPricesRes, err error)
	CreateTicketPrice(ctx context.Context, req *payment.PaymentCreateTicketPriceReq) (res *payment.PaymentCreateTicketPriceRes, err error)
	UpdateTicketPrice(ctx context.Context, req *payment.PaymentUpdateTicketPriceReq) (res *payment.PaymentUpdateTicketPriceRes, err error)
	DeleteTicketPrice(ctx context.Context, req *payment.PaymentDeleteTicketPriceReq) (res *payment.PaymentDeleteTicketPriceRes, err error)
}
