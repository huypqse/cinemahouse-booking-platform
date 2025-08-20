// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ScreeningSession is the golang structure for table screening_session.
type ScreeningSession struct {
	Id                     int64       `json:"id"                     orm:"id"                       description:""`
	StartDate              *gtime.Time `json:"startDate"              orm:"start_date"               description:""`
	StartTime              *gtime.Time `json:"startTime"              orm:"start_time"               description:""`
	EndTime                *gtime.Time `json:"endTime"                orm:"end_time"                 description:""`
	ScreeningSessionStatus string      `json:"screeningSessionStatus" orm:"screening_session_status" description:""`
	MovieId                int64       `json:"movieId"                orm:"movie_id"                 description:""`
	ScreeningRoomId        int64       `json:"screeningRoomId"        orm:"screening_room_id"        description:""`
}
