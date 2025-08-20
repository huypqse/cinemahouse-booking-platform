// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Ticket is the golang structure for table ticket.
type Ticket struct {
	Id                 int64  `json:"id"                 orm:"id"                   description:""`
	Status             string `json:"status"             orm:"status"               description:""`
	ScreeningSessionId int64  `json:"screeningSessionId" orm:"screening_session_id" description:""`
	TicketPriceId      int64  `json:"ticketPriceId"      orm:"ticket_price_id"      description:""`
	BillId             int64  `json:"billId"             orm:"bill_id"              description:""`
	SeatId             int64  `json:"seatId"             orm:"seat_id"              description:""`
}
