package cinema

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Room type management APIs
type CinemaGetRoomTypesReq struct {
	g.Meta `path:"/room-types" tags:"Cinema" method:"get" summary:"Get all room types"`
}

type CinemaGetRoomTypesRes struct {
	RoomTypes []RoomTypeInfo `json:"room_types"`
}

type RoomTypeInfo struct {
	RoomTypeId int64  `json:"room_type_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	RoomCount  int64  `json:"room_count"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CinemaCreateRoomTypeReq struct {
	g.Meta `path:"/room-types" tags:"Cinema" method:"post" summary:"Create new room type"`
	Name   string `json:"name" v:"required|length:1,100#Room type name is required|Name must be less than 100 characters"`
	Status string `json:"status" v:"required|in:ACTIVE,INACTIVE#Status is required|Status must be ACTIVE or INACTIVE"`
}

type CinemaCreateRoomTypeRes struct {
	RoomTypeId int64  `json:"room_type_id"`
	Message    string `json:"message"`
}

type CinemaUpdateRoomTypeReq struct {
	g.Meta     `path:"/room-types/{room_type_id}" tags:"Cinema" method:"put" summary:"Update room type"`
	RoomTypeId int64  `json:"room_type_id" v:"required|min:1#Room type ID is required|Room type ID must be positive"`
	Name       string `json:"name" v:"length:1,100#Name must be less than 100 characters"`
	Status     string `json:"status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
}

type CinemaUpdateRoomTypeRes struct {
	Message string `json:"message"`
}

type CinemaDeleteRoomTypeReq struct {
	g.Meta     `path:"/room-types/{room_type_id}" tags:"Cinema" method:"delete" summary:"Delete room type"`
	RoomTypeId int64 `json:"room_type_id" v:"required|min:1#Room type ID is required|Room type ID must be positive"`
}

type CinemaDeleteRoomTypeRes struct {
	Message string `json:"message"`
}
