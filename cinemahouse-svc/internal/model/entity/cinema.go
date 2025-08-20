// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Cinema is the golang structure for table cinema.
type Cinema struct {
	Id            int64  `json:"id"            orm:"id"             description:""`
	CinemaName    string `json:"cinemaName"    orm:"cinema_name"    description:""`
	CinemaAddress string `json:"cinemaAddress" orm:"cinema_address" description:""`
	CinemaStatus  string `json:"cinemaStatus"  orm:"cinema_status"  description:""`
}
