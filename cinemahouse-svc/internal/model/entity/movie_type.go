// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MovieType is the golang structure for table movie_type.
type MovieType struct {
	Id              int64  `json:"id"              orm:"id"                description:""`
	Type            string `json:"type"            orm:"type"              description:""`
	MovieTypeStatus string `json:"movieTypeStatus" orm:"movie_type_status" description:""`
}
