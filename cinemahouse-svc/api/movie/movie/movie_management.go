package movie

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Movie management APIs
type MovieGetMoviesReq struct {
	g.Meta    `path:"/movies" tags:"Movie" method:"get" summary:"Get movies with pagination and filters"`
	Page      int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize  int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	Search    string `json:"search"`
	Status    string `json:"status" v:"in:UPCOMING,RUNNING,ENDED#Status must be UPCOMING, RUNNING, or ENDED"`
	Language  string `json:"language"`
	AgeLimit  int    `json:"age_limit"`
	MovieType string `json:"movie_type"`
	StartDate string `json:"start_date" v:"date#Invalid start date format"`
	EndDate   string `json:"end_date" v:"date#Invalid end date format"`
}

type MovieGetMoviesRes struct {
	Movies     []MovieInfo `json:"movies"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

type MovieInfo struct {
	MovieId     int64    `json:"movie_id"`
	MovieName   string   `json:"movie_name"`
	Content     string   `json:"content"`
	Duration    string   `json:"duration"`
	Language    string   `json:"language"`
	Director    string   `json:"director"`
	Actors      string   `json:"actors"`
	AgeLimit    int      `json:"age_limit"`
	CoverPhoto  string   `json:"cover_photo"`
	ReleaseDate string   `json:"release_date"`
	StartDate   string   `json:"start_date"`
	Status      string   `json:"status"`
	MovieTypes  []string `json:"movie_types"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type MovieGetMovieReq struct {
	g.Meta  `path:"/movies/{movie_id}" tags:"Movie" method:"get" summary:"Get movie by ID"`
	MovieId int64 `json:"movie_id" v:"required|min:1#Movie ID is required|Movie ID must be positive"`
}

type MovieGetMovieRes struct {
	Movie MovieInfo `json:"movie"`
}

type MovieCreateMovieReq struct {
	g.Meta      `path:"/movies" tags:"Movie" method:"post" summary:"Create new movie"`
	MovieName   string  `json:"movie_name" v:"required|length:1,255#Movie name is required|Movie name must be less than 255 characters"`
	Content     string  `json:"content" v:"required#Movie content is required"`
	Duration    string  `json:"duration" v:"required#Duration is required"`
	Language    string  `json:"language" v:"required|length:1,100#Language is required|Language must be less than 100 characters"`
	Director    string  `json:"director" v:"required|length:1,255#Director is required|Director must be less than 255 characters"`
	Actors      string  `json:"actors" v:"required|length:1,255#Actors are required|Actors must be less than 255 characters"`
	AgeLimit    int     `json:"age_limit" v:"required|min:0|max:21#Age limit is required|Age limit must be between 0 and 21"`
	CoverPhoto  string  `json:"cover_photo" v:"required#Cover photo is required"`
	ReleaseDate string  `json:"release_date" v:"required|date#Release date is required|Invalid release date format"`
	StartDate   string  `json:"start_date" v:"required|date#Start date is required|Invalid start date format"`
	Status      string  `json:"status" v:"required|in:UPCOMING,RUNNING,ENDED#Status is required|Status must be UPCOMING, RUNNING, or ENDED"`
	MovieTypes  []int64 `json:"movie_types"`
}

type MovieCreateMovieRes struct {
	MovieId int64  `json:"movie_id"`
	Message string `json:"message"`
}

type MovieUpdateMovieReq struct {
	g.Meta      `path:"/movies/{movie_id}" tags:"Movie" method:"put" summary:"Update movie"`
	MovieId     int64   `json:"movie_id" v:"required|min:1#Movie ID is required|Movie ID must be positive"`
	MovieName   string  `json:"movie_name" v:"length:1,255#Movie name must be less than 255 characters"`
	Content     string  `json:"content"`
	Duration    string  `json:"duration"`
	Language    string  `json:"language" v:"length:1,100#Language must be less than 100 characters"`
	Director    string  `json:"director" v:"length:1,255#Director must be less than 255 characters"`
	Actors      string  `json:"actors" v:"length:1,255#Actors must be less than 255 characters"`
	AgeLimit    int     `json:"age_limit" v:"min:0|max:21#Age limit must be between 0 and 21"`
	CoverPhoto  string  `json:"cover_photo"`
	ReleaseDate string  `json:"release_date" v:"date#Invalid release date format"`
	StartDate   string  `json:"start_date" v:"date#Invalid start date format"`
	Status      string  `json:"status" v:"in:UPCOMING,RUNNING,ENDED#Status must be UPCOMING, RUNNING, or ENDED"`
	MovieTypes  []int64 `json:"movie_types"`
}

type MovieUpdateMovieRes struct {
	Message string `json:"message"`
}

type MovieDeleteMovieReq struct {
	g.Meta  `path:"/movies/{movie_id}" tags:"Movie" method:"delete" summary:"Delete movie"`
	MovieId int64 `json:"movie_id" v:"required|min:1#Movie ID is required|Movie ID must be positive"`
}

type MovieDeleteMovieRes struct {
	Message string `json:"message"`
}
