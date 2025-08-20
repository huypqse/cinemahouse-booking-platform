package movie

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Movie type management APIs
type MovieGetMovieTypesReq struct {
	g.Meta `path:"/movie-types" tags:"Movie" method:"get" summary:"Get all movie types"`
}

type MovieGetMovieTypesRes struct {
	MovieTypes []MovieTypeInfo `json:"movie_types"`
}

type MovieTypeInfo struct {
	MovieTypeId     int64  `json:"movie_type_id"`
	Type            string `json:"type"`
	MovieTypeStatus string `json:"movie_type_status"`
	MovieCount      int64  `json:"movie_count"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type MovieCreateMovieTypeReq struct {
	g.Meta          `path:"/movie-types" tags:"Movie" method:"post" summary:"Create new movie type"`
	Type            string `json:"type" v:"required|length:1,100#Movie type is required|Movie type must be less than 100 characters"`
	MovieTypeStatus string `json:"movie_type_status" v:"required|in:ACTIVE,INACTIVE#Movie type status is required|Status must be ACTIVE or INACTIVE"`
}

type MovieCreateMovieTypeRes struct {
	MovieTypeId int64  `json:"movie_type_id"`
	Message     string `json:"message"`
}

type MovieUpdateMovieTypeReq struct {
	g.Meta          `path:"/movie-types/{movie_type_id}" tags:"Movie" method:"put" summary:"Update movie type"`
	MovieTypeId     int64  `json:"movie_type_id" v:"required|min:1#Movie type ID is required|Movie type ID must be positive"`
	Type            string `json:"type" v:"length:1,100#Movie type must be less than 100 characters"`
	MovieTypeStatus string `json:"movie_type_status" v:"in:ACTIVE,INACTIVE#Status must be ACTIVE or INACTIVE"`
}

type MovieUpdateMovieTypeRes struct {
	Message string `json:"message"`
}

type MovieDeleteMovieTypeReq struct {
	g.Meta      `path:"/movie-types/{movie_type_id}" tags:"Movie" method:"delete" summary:"Delete movie type"`
	MovieTypeId int64 `json:"movie_type_id" v:"required|min:1#Movie type ID is required|Movie type ID must be positive"`
}

type MovieDeleteMovieTypeRes struct {
	Message string `json:"message"`
}
