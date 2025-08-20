package booking

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Seat management APIs
type BookingGetSeatsReq struct {
	g.Meta          `path:"/seats" tags:"Booking" method:"get" summary:"Get seats with filters"`
	Page            int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize        int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	ScreeningRoomId int64  `json:"screening_room_id"`
	SeatTypeId      int64  `json:"seat_type_id"`
	Status          string `json:"status" v:"in:AVAILABLE,RESERVED,SOLD#Status must be AVAILABLE, RESERVED, or SOLD"`
}

type BookingGetSeatsRes struct {
	Seats      []SeatInfo `json:"seats"`
	Total      int64      `json:"total"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}

type SeatInfo struct {
	SeatId            int64  `json:"seat_id"`
	SeatRow           string `json:"seat_row"`
	SeatColumn        int    `json:"seat_column"`
	Status            string `json:"status"`
	ScreeningRoomName string `json:"screening_room_name"`
	SeatTypeName      string `json:"seat_type_name"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type BookingGetSeatsBySessionReq struct {
	g.Meta             `path:"/screening-sessions/{screening_session_id}/seats" tags:"Booking" method:"get" summary:"Get seats for a screening session"`
	ScreeningSessionId int64 `json:"screening_session_id" v:"required|min:1#Screening session ID is required|Screening session ID must be positive"`
}

type BookingGetSeatsBySessionRes struct {
	Seats []SeatSessionInfo `json:"seats"`
}

type SeatSessionInfo struct {
	SeatId       int64   `json:"seat_id"`
	SeatRow      string  `json:"seat_row"`
	SeatColumn   int     `json:"seat_column"`
	Status       string  `json:"status"`
	SeatTypeName string  `json:"seat_type_name"`
	Price        float64 `json:"price"`
	IsAvailable  bool    `json:"is_available"`
}

type BookingCreateSeatReq struct {
	g.Meta          `path:"/seats" tags:"Booking" method:"post" summary:"Create new seat"`
	SeatRow         string `json:"seat_row" v:"required|length:1,10#Seat row is required|Seat row must be less than 10 characters"`
	SeatColumn      int    `json:"seat_column" v:"required|min:1#Seat column is required|Seat column must be positive"`
	Status          string `json:"status" v:"required|in:AVAILABLE,RESERVED,SOLD#Status is required|Status must be AVAILABLE, RESERVED, or SOLD"`
	ScreeningRoomId int64  `json:"screening_room_id" v:"required|min:1#Screening room ID is required|Screening room ID must be positive"`
	SeatTypeId      int64  `json:"seat_type_id" v:"required|min:1#Seat type ID is required|Seat type ID must be positive"`
}

type BookingCreateSeatRes struct {
	SeatId  int64  `json:"seat_id"`
	Message string `json:"message"`
}

type BookingUpdateSeatReq struct {
	g.Meta          `path:"/seats/{seat_id}" tags:"Booking" method:"put" summary:"Update seat"`
	SeatId          int64  `json:"seat_id" v:"required|min:1#Seat ID is required|Seat ID must be positive"`
	SeatRow         string `json:"seat_row" v:"length:1,10#Seat row must be less than 10 characters"`
	SeatColumn      int    `json:"seat_column" v:"min:1#Seat column must be positive"`
	Status          string `json:"status" v:"in:AVAILABLE,RESERVED,SOLD#Status must be AVAILABLE, RESERVED, or SOLD"`
	ScreeningRoomId int64  `json:"screening_room_id" v:"min:1#Screening room ID must be positive"`
	SeatTypeId      int64  `json:"seat_type_id" v:"min:1#Seat type ID must be positive"`
}

type BookingUpdateSeatRes struct {
	Message string `json:"message"`
}

type BookingDeleteSeatReq struct {
	g.Meta `path:"/seats/{seat_id}" tags:"Booking" method:"delete" summary:"Delete seat"`
	SeatId int64 `json:"seat_id" v:"required|min:1#Seat ID is required|Seat ID must be positive"`
}

type BookingDeleteSeatRes struct {
	Message string `json:"message"`
}
