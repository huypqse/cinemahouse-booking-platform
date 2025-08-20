package payment

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Ticket price management APIs
type PaymentGetTicketPricesReq struct {
	g.Meta `path:"/ticket-prices" tags:"Payment" method:"get" summary:"Get all ticket prices"`
}

type PaymentGetTicketPricesRes struct {
	TicketPrices []TicketPriceInfo `json:"ticket_prices"`
}

type TicketPriceInfo struct {
	TicketPriceId int64   `json:"ticket_price_id"`
	Price         float64 `json:"price"`
	Status        string  `json:"status"`
	Description   string  `json:"description"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type PaymentCreateTicketPriceReq struct {
	g.Meta      `path:"/ticket-prices" tags:"Payment" method:"post" summary:"Create new ticket price"`
	Price       float64 `json:"price" v:"required|min:0#Price is required|Price must be positive"`
	Status      string  `json:"status" v:"required|in:ACTIVE,INACTIVE#Status is required|Status must be ACTIVE or INACTIVE"`
	Description string  `json:"description"`
}

type PaymentCreateTicketPriceRes struct {
	TicketPriceId int64  `json:"ticket_price_id"`
	Message       string `json:"message"`
}

type PaymentUpdateTicketPriceReq struct {
	g.Meta        `path:"/ticket-prices/{ticket_price_id}" tags:"Payment" method:"put" summary:"Update ticket price"`
	TicketPriceId int64   `json:"ticket_price_id" v:"required|min:1#Ticket price ID is required|Ticket price ID must be positive"`
	Price         float64 `json:"price" v:"min:0#Price must be positive"`
	Status        string  `json:"status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
	Description   string  `json:"description"`
}

type PaymentUpdateTicketPriceRes struct {
	Message string `json:"message"`
}

type PaymentDeleteTicketPriceReq struct {
	g.Meta        `path:"/ticket-prices/{ticket_price_id}" tags:"Payment" method:"delete" summary:"Delete ticket price"`
	TicketPriceId int64 `json:"ticket_price_id" v:"required|min:1#Ticket price ID is required|Ticket price ID must be positive"`
}

type PaymentDeleteTicketPriceRes struct {
	Message string `json:"message"`
}
