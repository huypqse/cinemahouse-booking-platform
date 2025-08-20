package booking

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Bill management APIs
type BookingGetBillsReq struct {
	g.Meta    `path:"/bills" tags:"Booking" method:"get" summary:"Get bills with filters"`
	Page      int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize  int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	UserId    int64  `json:"user_id"`
	Status    string `json:"status" v:"in:PENDING,PAID,CANCELLED#Status must be PENDING, PAID, or CANCELLED"`
	StartDate string `json:"start_date" v:"date#Invalid start date format"`
	EndDate   string `json:"end_date" v:"date#Invalid end date format"`
}

type BookingGetBillsRes struct {
	Bills      []BillInfo `json:"bills"`
	Total      int64      `json:"total"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}

type BillInfo struct {
	BillId      int64        `json:"bill_id"`
	BillDate    string       `json:"bill_date"`
	BillTime    string       `json:"bill_time"`
	TotalAmount float64      `json:"total_amount"`
	Status      string       `json:"status"`
	UserId      int64        `json:"user_id"`
	PaymentId   int64        `json:"payment_id,omitempty"`
	RequestId   string       `json:"request_id,omitempty"`
	TicketCount int          `json:"ticket_count"`
	Tickets     []TicketInfo `json:"tickets,omitempty"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type BookingGetBillReq struct {
	g.Meta `path:"/bills/{bill_id}" tags:"Booking" method:"get" summary:"Get bill by ID"`
	BillId int64 `json:"bill_id" v:"required|min:1#Bill ID is required|Bill ID must be positive"`
}

type BookingGetBillRes struct {
	Bill BillInfo `json:"bill"`
}

type BookingUpdateBillStatusReq struct {
	g.Meta `path:"/bills/{bill_id}/status" tags:"Booking" method:"put" summary:"Update bill status"`
	BillId int64  `json:"bill_id" v:"required|min:1#Bill ID is required|Bill ID must be positive"`
	Status string `json:"status" v:"required|in:PENDING,PAID,CANCELLED#Status is required|Status must be PENDING, PAID, or CANCELLED"`
}

type BookingUpdateBillStatusRes struct {
	Message string `json:"message"`
}
