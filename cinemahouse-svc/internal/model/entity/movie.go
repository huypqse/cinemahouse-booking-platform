// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Movie is the golang structure for table movie.
type Movie struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	MovieName   string      `json:"movieName"   orm:"movie_name"   description:""`
	Content     string      `json:"content"     orm:"content"      description:""`
	Duration    *gtime.Time `json:"duration"    orm:"duration"     description:""`
	Language    string      `json:"language"    orm:"language"     description:""`
	Director    string      `json:"director"    orm:"director"     description:""`
	Actors      string      `json:"actors"      orm:"actors"       description:""`
	AgeLimit    uint        `json:"ageLimit"    orm:"age_limit"    description:""`
	CoverPhoto  string      `json:"coverPhoto"  orm:"cover_photo"  description:""`
	ReleaseDate *gtime.Time `json:"releaseDate" orm:"release_date" description:""`
	StartDate   *gtime.Time `json:"startDate"   orm:"start_date"   description:""`
	Status      string      `json:"status"      orm:"status"       description:""`
}
