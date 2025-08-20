// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Comment is the golang structure for table comment.
type Comment struct {
	Id      int64  `json:"id"      orm:"id"       description:""`
	Content string `json:"content" orm:"content"  description:""`
	Status  string `json:"status"  orm:"status"   description:""`
	MovieId int64  `json:"movieId" orm:"movie_id" description:""`
}
