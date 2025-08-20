package cinema

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Cinema management APIs
type CinemaGetCinemasReq struct {
	g.Meta   `path:"/cinemas" tags:"Cinema" method:"get" summary:"Get cinemas with pagination and filters"`
	Page     int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	Search   string `json:"search"`
	Status   string `json:"status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
	City     string `json:"city"`
}

type CinemaGetCinemasRes struct {
	Cinemas    []CinemaInfo `json:"cinemas"`
	Total      int64        `json:"total"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
	TotalPages int          `json:"total_pages"`
}

type CinemaInfo struct {
	CinemaId       int64  `json:"cinema_id"`
	CinemaName     string `json:"cinema_name"`
	CinemaAddress  string `json:"cinema_address"`
	CinemaStatus   string `json:"cinema_status"`
	ScreeningRooms int    `json:"screening_rooms"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CinemaGetCinemaReq struct {
	g.Meta   `path:"/cinemas/{cinema_id}" tags:"Cinema" method:"get" summary:"Get cinema by ID"`
	CinemaId int64 `json:"cinema_id" v:"required|min:1#Cinema ID is required|Cinema ID must be positive"`
}

type CinemaGetCinemaRes struct {
	Cinema CinemaInfo `json:"cinema"`
}

type CinemaCreateCinemaReq struct {
	g.Meta        `path:"/cinemas" tags:"Cinema" method:"post" summary:"Create new cinema"`
	CinemaName    string `json:"cinema_name" v:"required|length:1,255#Cinema name is required|Cinema name must be less than 255 characters"`
	CinemaAddress string `json:"cinema_address" v:"required|length:1,255#Cinema address is required|Cinema address must be less than 255 characters"`
	CinemaStatus  string `json:"cinema_status" v:"required|in:ACTIVE,INACTIVE#Cinema status is required|Status must be ACTIVE or INACTIVE"`
}

type CinemaCreateCinemaRes struct {
	CinemaId int64  `json:"cinema_id"`
	Message  string `json:"message"`
}

type CinemaUpdateCinemaReq struct {
	g.Meta        `path:"/cinemas/{cinema_id}" tags:"Cinema" method:"put" summary:"Update cinema"`
	CinemaId      int64  `json:"cinema_id" v:"required|min:1#Cinema ID is required|Cinema ID must be positive"`
	CinemaName    string `json:"cinema_name" v:"length:1,255#Cinema name must be less than 255 characters"`
	CinemaAddress string `json:"cinema_address" v:"length:1,255#Cinema address must be less than 255 characters"`
	CinemaStatus  string `json:"cinema_status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
}

type CinemaUpdateCinemaRes struct {
	Message string `json:"message"`
}

type CinemaDeleteCinemaReq struct {
	g.Meta   `path:"/cinemas/{cinema_id}" tags:"Cinema" method:"delete" summary:"Delete cinema"`
	CinemaId int64 `json:"cinema_id" v:"required|min:1#Cinema ID is required|Cinema ID must be positive"`
}

type CinemaDeleteCinemaRes struct {
	Message string `json:"message"`
}
