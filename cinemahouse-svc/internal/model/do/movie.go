// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Movie is the golang structure of table movie for DAO operations like Where/Data.
type Movie struct {
	g.Meta      `orm:"table:movie, do:true"`
	Id          interface{} //
	MovieName   interface{} //
	Content     interface{} //
	Duration    *gtime.Time //
	Language    interface{} //
	Director    interface{} //
	Actors      interface{} //
	AgeLimit    interface{} //
	CoverPhoto  interface{} //
	ReleaseDate *gtime.Time //
	StartDate   *gtime.Time //
	Status      interface{} //
}
