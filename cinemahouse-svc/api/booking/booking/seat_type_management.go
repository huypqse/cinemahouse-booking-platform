package booking

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Seat type management APIs
type BookingSeatTypesReq struct {
	g.Meta `path:"/seat-types" tags:"Booking" method:"get" summary:"Get all seat types"`
}

type BookingSeatTypesRes struct {
	SeatTypes []SeatTypeInfo `json:"seat_types"`
}

type SeatTypeInfo struct {
	SeatTypeId int64  `json:"seat_type_id"`
	Name       string `json:"name"`
	SeatCount  int64  `json:"seat_count"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type BookingCreateSeatTypeReq struct {
	g.Meta `path:"/seat-types" tags:"Booking" method:"post" summary:"Create new seat type"`
	Name   string `json:"name" v:"required|length:1,100#Seat type name is required|Name must be less than 100 characters"`
}

type BookingCreateSeatTypeRes struct {
	SeatTypeId int64  `json:"seat_type_id"`
	Message    string `json:"message"`
}

type BookingUpdateSeatTypeReq struct {
	g.Meta     `path:"/seat-types/{seat_type_id}" tags:"Booking" method:"put" summary:"Update seat type"`
	SeatTypeId int64  `json:"seat_type_id" v:"required|min:1#Seat type ID is required|Seat type ID must be positive"`
	Name       string `json:"name" v:"required|length:1,100#Seat type name is required|Name must be less than 100 characters"`
}

type BookingUpdateSeatTypeRes struct {
	Message string `json:"message"`
}

type BookingDeleteSeatTypeReq struct {
	g.Meta     `path:"/seat-types/{seat_type_id}" tags:"Booking" method:"delete" summary:"Delete seat type"`
	SeatTypeId int64 `json:"seat_type_id" v:"required|min:1#Seat type ID is required|Seat type ID must be positive"`
}

type BookingDeleteSeatTypeRes struct {
	Message string `json:"message"`
}
