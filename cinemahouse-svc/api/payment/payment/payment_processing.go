package payment

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Payment processing APIs
type PaymentCreatePaymentReq struct {
	g.Meta        `path:"/payments" tags:"Payment" method:"post" summary:"Create payment"`
	BillId        int64   `json:"bill_id" v:"required|min:1#Bill ID is required|Bill ID must be positive"`
	PaymentMethod string  `json:"payment_method" v:"required|in:CREDIT_CARD,DEBIT_CARD,PAYPAL,BANK_TRANSFER,CASH#Payment method is required|Invalid payment method"`
	Amount        float64 `json:"amount" v:"required|min:0#Amount is required|Amount must be positive"`
}

type PaymentCreatePaymentRes struct {
	PaymentId     int64  `json:"payment_id"`
	TransactionId string `json:"transaction_id"`
	PaymentUrl    string `json:"payment_url,omitempty"`
	Message       string `json:"message"`
}

type PaymentProcessPaymentReq struct {
	g.Meta        `path:"/payments/{payment_id}/process" tags:"Payment" method:"post" summary:"Process payment"`
	PaymentId     int64  `json:"payment_id" v:"required|min:1#Payment ID is required|Payment ID must be positive"`
	TransactionId string `json:"transaction_id" v:"required#Transaction ID is required"`
	CardNumber    string `json:"card_number,omitempty"`
	ExpiryDate    string `json:"expiry_date,omitempty"`
	CVV           string `json:"cvv,omitempty"`
	CardHolder    string `json:"card_holder,omitempty"`
}

type PaymentProcessPaymentRes struct {
	PaymentStatus string `json:"payment_status"`
	TransactionId string `json:"transaction_id"`
	Message       string `json:"message"`
}

type PaymentGetPaymentReq struct {
	g.Meta    `path:"/payments/{payment_id}" tags:"Payment" method:"get" summary:"Get payment by ID"`
	PaymentId int64 `json:"payment_id" v:"required|min:1#Payment ID is required|Payment ID must be positive"`
}

type PaymentGetPaymentRes struct {
	Payment PaymentInfo `json:"payment"`
}

type PaymentInfo struct {
	PaymentId     int64   `json:"payment_id"`
	PaymentMethod string  `json:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
	PaymentDate   string  `json:"payment_date"`
	Amount        float64 `json:"amount"`
	TransactionId string  `json:"transaction_id"`
	BillId        int64   `json:"bill_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type PaymentGetPaymentsReq struct {
	g.Meta        `path:"/payments" tags:"Payment" method:"get" summary:"Get payments with filters"`
	Page          int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize      int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	PaymentMethod string `json:"payment_method" v:"in:CREDIT_CARD,DEBIT_CARD,PAYPAL,BANK_TRANSFER,CASH#Invalid payment method"`
	PaymentStatus string `json:"payment_status" v:"in:PENDING,COMPLETED,FAILED,CANCELLED#Invalid payment status"`
	StartDate     string `json:"start_date" v:"date#Invalid start date format"`
	EndDate       string `json:"end_date" v:"date#Invalid end date format"`
}

type PaymentGetPaymentsRes struct {
	Payments   []PaymentInfo `json:"payments"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
	TotalPages int           `json:"total_pages"`
}

type PaymentRefundPaymentReq struct {
	g.Meta    `path:"/payments/{payment_id}/refund" tags:"Payment" method:"post" summary:"Refund payment"`
	PaymentId int64   `json:"payment_id" v:"required|min:1#Payment ID is required|Payment ID must be positive"`
	Amount    float64 `json:"amount" v:"required|min:0#Refund amount is required|Amount must be positive"`
	Reason    string  `json:"reason" v:"required#Refund reason is required"`
}

type PaymentRefundPaymentRes struct {
	RefundId      int64  `json:"refund_id"`
	TransactionId string `json:"transaction_id"`
	Message       string `json:"message"`
}
