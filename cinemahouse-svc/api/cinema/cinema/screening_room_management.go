package cinema

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Screening room management APIs
type CinemaGetScreeningRoomsReq struct {
	g.Meta   `path:"/screening-rooms" tags:"Cinema" method:"get" summary:"Get screening rooms with pagination and filters"`
	Page     int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	CinemaId int64  `json:"cinema_id"`
	Status   string `json:"status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
	RoomType string `json:"room_type"`
}

type CinemaGetScreeningRoomsRes struct {
	ScreeningRooms []ScreeningRoomInfo `json:"screening_rooms"`
	Total          int64               `json:"total"`
	Page           int                 `json:"page"`
	PageSize       int                 `json:"page_size"`
	TotalPages     int                 `json:"total_pages"`
}

type ScreeningRoomInfo struct {
	ScreeningRoomId     int64  `json:"screening_room_id"`
	Name                string `json:"name"`
	ScreeningRoomStatus string `json:"screening_room_status"`
	RoomTypeName        string `json:"room_type_name"`
	CinemaName          string `json:"cinema_name"`
	SeatCount           int    `json:"seat_count"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

type CinemaGetScreeningRoomReq struct {
	g.Meta          `path:"/screening-rooms/{screening_room_id}" tags:"Cinema" method:"get" summary:"Get screening room by ID"`
	ScreeningRoomId int64 `json:"screening_room_id" v:"required|min:1#Screening room ID is required|Screening room ID must be positive"`
}

type CinemaGetScreeningRoomRes struct {
	ScreeningRoom ScreeningRoomInfo `json:"screening_room"`
}

type CinemaCreateScreeningRoomReq struct {
	g.Meta              `path:"/screening-rooms" tags:"Cinema" method:"post" summary:"Create new screening room"`
	Name                string `json:"name" v:"required|length:1,100#Screening room name is required|Name must be less than 100 characters"`
	ScreeningRoomStatus string `json:"screening_room_status" v:"required|in:ACTIVE,INACTIVE#Status is required|Status must be ACTIVE or INACTIVE"`
	RoomTypeId          int64  `json:"room_type_id" v:"required|min:1#Room type ID is required|Room type ID must be positive"`
	CinemaId            int64  `json:"cinema_id" v:"required|min:1#Cinema ID is required|Cinema ID must be positive"`
}

type CinemaCreateScreeningRoomRes struct {
	ScreeningRoomId int64  `json:"screening_room_id"`
	Message         string `json:"message"`
}

type CinemaUpdateScreeningRoomReq struct {
	g.Meta              `path:"/screening-rooms/{screening_room_id}" tags:"Cinema" method:"put" summary:"Update screening room"`
	ScreeningRoomId     int64  `json:"screening_room_id" v:"required|min:1#Screening room ID is required|Screening room ID must be positive"`
	Name                string `json:"name" v:"length:1,100#Name must be less than 100 characters"`
	ScreeningRoomStatus string `json:"screening_room_status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
	RoomTypeId          int64  `json:"room_type_id" v:"min:1#Room type ID must be positive"`
	CinemaId            int64  `json:"cinema_id" v:"min:1#Cinema ID must be positive"`
}

type CinemaUpdateScreeningRoomRes struct {
	Message string `json:"message"`
}

type CinemaDeleteScreeningRoomReq struct {
	g.Meta          `path:"/screening-rooms/{screening_room_id}" tags:"Cinema" method:"delete" summary:"Delete screening room"`
	ScreeningRoomId int64 `json:"screening_room_id" v:"required|min:1#Screening room ID is required|Screening room ID must be positive"`
}

type CinemaDeleteScreeningRoomRes struct {
	Message string `json:"message"`
}
