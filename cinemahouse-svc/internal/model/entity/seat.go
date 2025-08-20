// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Seat is the golang structure for table seat.
type Seat struct {
	Id              int64  `json:"id"              orm:"id"                description:""`
	SeatRow         string `json:"seatRow"         orm:"seat_row"          description:""`
	SeatColumn      uint   `json:"seatColumn"      orm:"seat_column"       description:""`
	Status          string `json:"status"          orm:"status"            description:""`
	ScreeningRoomId int64  `json:"screeningRoomId" orm:"screening_room_id" description:""`
	SeatTypeId      int64  `json:"seatTypeId"      orm:"seat_type_id"      description:""`
}
