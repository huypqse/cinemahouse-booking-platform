// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomType is the golang structure for table room_type.
type RoomType struct {
	Id     int64  `json:"id"     orm:"id"     description:""`
	Name   string `json:"name"   orm:"name"   description:""`
	Status string `json:"status" orm:"status" description:""`
}
