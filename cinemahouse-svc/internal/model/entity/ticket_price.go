// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TicketPrice is the golang structure for table ticket_price.
type TicketPrice struct {
	Id     int64   `json:"id"     orm:"id"     description:""`
	Price  float64 `json:"price"  orm:"price"  description:""`
	Status string  `json:"status" orm:"status" description:""`
}
