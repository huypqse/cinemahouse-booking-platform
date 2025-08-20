// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ScreeningRoom is the golang structure for table screening_room.
type ScreeningRoom struct {
	Id                  int64  `json:"id"                  orm:"id"                    description:""`
	Name                string `json:"name"                orm:"name"                  description:""`
	ScreeningRoomStatus string `json:"screeningRoomStatus" orm:"screening_room_status" description:""`
	RoomtypeId          int64  `json:"roomtypeId"          orm:"roomtype_id"           description:""`
	CinemaId            int64  `json:"cinemaId"            orm:"cinema_id"             description:""`
}
