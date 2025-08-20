package booking

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Screening session management APIs
type BookingGetScreeningSessionsReq struct {
	g.Meta          `path:"/screening-sessions" tags:"Booking" method:"get" summary:"Get screening sessions with filters"`
	Page            int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize        int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	MovieId         int64  `json:"movie_id"`
	CinemaId        int64  `json:"cinema_id"`
	ScreeningRoomId int64  `json:"screening_room_id"`
	StartDate       string `json:"start_date" v:"date#Invalid start date format"`
	EndDate         string `json:"end_date" v:"date#Invalid end date format"`
	Status          string `json:"status" v:"in:SCHEDULED,ONGOING,ENDED#Status must be SCHEDULED, ONGOING, or ENDED"`
}

type BookingGetScreeningSessionsRes struct {
	ScreeningSessions []ScreeningSessionInfo `json:"screening_sessions"`
	Total             int64                  `json:"total"`
	Page              int                    `json:"page"`
	PageSize          int                    `json:"page_size"`
	TotalPages        int                    `json:"total_pages"`
}

type ScreeningSessionInfo struct {
	ScreeningSessionId     int64  `json:"screening_session_id"`
	StartDate              string `json:"start_date"`
	StartTime              string `json:"start_time"`
	EndTime                string `json:"end_time"`
	ScreeningSessionStatus string `json:"screening_session_status"`
	MovieName              string `json:"movie_name"`
	CinemaName             string `json:"cinema_name"`
	ScreeningRoomName      string `json:"screening_room_name"`
	AvailableSeats         int    `json:"available_seats"`
	TotalSeats             int    `json:"total_seats"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

type BookingGetScreeningSessionReq struct {
	g.Meta             `path:"/screening-sessions/{screening_session_id}" tags:"Booking" method:"get" summary:"Get screening session by ID"`
	ScreeningSessionId int64 `json:"screening_session_id" v:"required|min:1#Screening session ID is required|Screening session ID must be positive"`
}

type BookingGetScreeningSessionRes struct {
	ScreeningSession ScreeningSessionInfo `json:"screening_session"`
}

type BookingCreateScreeningSessionReq struct {
	g.Meta                 `path:"/screening-sessions" tags:"Booking" method:"post" summary:"Create new screening session"`
	StartDate              string `json:"start_date" v:"required|date#Start date is required|Invalid start date format"`
	StartTime              string `json:"start_time" v:"required#Start time is required"`
	EndTime                string `json:"end_time" v:"required#End time is required"`
	ScreeningSessionStatus string `json:"screening_session_status" v:"required|in:SCHEDULED,ONGOING,ENDED#Status is required|Status must be SCHEDULED, ONGOING, or ENDED"`
	MovieId                int64  `json:"movie_id" v:"required|min:1#Movie ID is required|Movie ID must be positive"`
	ScreeningRoomId        int64  `json:"screening_room_id" v:"required|min:1#Screening room ID is required|Screening room ID must be positive"`
}

type BookingCreateScreeningSessionRes struct {
	ScreeningSessionId int64  `json:"screening_session_id"`
	Message            string `json:"message"`
}

type BookingUpdateScreeningSessionReq struct {
	g.Meta                 `path:"/screening-sessions/{screening_session_id}" tags:"Booking" method:"put" summary:"Update screening session"`
	ScreeningSessionId     int64  `json:"screening_session_id" v:"required|min:1#Screening session ID is required|Screening session ID must be positive"`
	StartDate              string `json:"start_date" v:"date#Invalid start date format"`
	StartTime              string `json:"start_time"`
	EndTime                string `json:"end_time"`
	ScreeningSessionStatus string `json:"screening_session_status" v:"in:SCHEDULED,ONGOING,ENDED#Status must be SCHEDULED, ONGOING, or ENDED"`
	MovieId                int64  `json:"movie_id" v:"min:1#Movie ID must be positive"`
	ScreeningRoomId        int64  `json:"screening_room_id" v:"min:1#Screening room ID must be positive"`
}

type BookingUpdateScreeningSessionRes struct {
	Message string `json:"message"`
}

type BookingDeleteScreeningSessionReq struct {
	g.Meta             `path:"/screening-sessions/{screening_session_id}" tags:"Booking" method:"delete" summary:"Delete screening session"`
	ScreeningSessionId int64 `json:"screening_session_id" v:"required|min:1#Screening session ID is required|Screening session ID must be positive"`
}

type BookingDeleteScreeningSessionRes struct {
	Message string `json:"message"`
}
