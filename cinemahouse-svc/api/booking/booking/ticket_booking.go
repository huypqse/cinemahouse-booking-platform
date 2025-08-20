package booking

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Ticket booking APIs
type BookingBookTicketsReq struct {
	g.Meta             `path:"/tickets/book" tags:"Booking" method:"post" summary:"Book tickets"`
	UserId             int64   `json:"user_id" v:"required|min:1#User ID is required|User ID must be positive"`
	ScreeningSessionId int64   `json:"screening_session_id" v:"required|min:1#Screening session ID is required|Screening session ID must be positive"`
	SeatIds            []int64 `json:"seat_ids" v:"required#Seat IDs are required"`
	PaymentMethod      string  `json:"payment_method" v:"required|in:CREDIT_CARD,DEBIT_CARD,PAYPAL,BANK_TRANSFER,CASH#Payment method is required|Invalid payment method"`
}

type BookingBookTicketsRes struct {
	BillId    int64        `json:"bill_id"`
	TicketIds []int64      `json:"ticket_ids"`
	Total     float64      `json:"total"`
	Tickets   []TicketInfo `json:"tickets"`
	Message   string       `json:"message"`
}

type TicketInfo struct {
	TicketId          int64   `json:"ticket_id"`
	Status            string  `json:"status"`
	SeatRow           string  `json:"seat_row"`
	SeatColumn        int     `json:"seat_column"`
	Price             float64 `json:"price"`
	MovieName         string  `json:"movie_name"`
	CinemaName        string  `json:"cinema_name"`
	ScreeningRoomName string  `json:"screening_room_name"`
	ScreeningDate     string  `json:"screening_date"`
	ScreeningTime     string  `json:"screening_time"`
	CreatedAt         string  `json:"created_at"`
}

type BookingGetTicketsReq struct {
	g.Meta   `path:"/tickets" tags:"Booking" method:"get" summary:"Get tickets with filters"`
	Page     int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	UserId   int64  `json:"user_id"`
	Status   string `json:"status" v:"in:AVAILABLE,SOLD,CANCELLED#Status must be AVAILABLE, SOLD, or CANCELLED"`
	MovieId  int64  `json:"movie_id"`
	CinemaId int64  `json:"cinema_id"`
}

type BookingGetTicketsRes struct {
	Tickets    []TicketInfo `json:"tickets"`
	Total      int64        `json:"total"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
	TotalPages int          `json:"total_pages"`
}

type BookingGetTicketReq struct {
	g.Meta   `path:"/tickets/{ticket_id}" tags:"Booking" method:"get" summary:"Get ticket by ID"`
	TicketId int64 `json:"ticket_id" v:"required|min:1#Ticket ID is required|Ticket ID must be positive"`
}

type BookingGetTicketRes struct {
	Ticket TicketInfo `json:"ticket"`
}

type BookingCancelTicketReq struct {
	g.Meta   `path:"/tickets/{ticket_id}/cancel" tags:"Booking" method:"post" summary:"Cancel ticket"`
	TicketId int64  `json:"ticket_id" v:"required|min:1#Ticket ID is required|Ticket ID must be positive"`
	Reason   string `json:"reason" v:"required#Cancellation reason is required"`
}

type BookingCancelTicketRes struct {
	Message string `json:"message"`
}
