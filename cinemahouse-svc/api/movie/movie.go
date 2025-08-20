// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package movie

import (
	"context"

	"cinemahouse-svc/api/movie/movie"
)

type IMovieV1 interface {
	// Movie management
	GetMovies(ctx context.Context, req *movie.MovieGetMoviesReq) (res *movie.MovieGetMoviesRes, err error)
	GetMovie(ctx context.Context, req *movie.MovieGetMovieReq) (res *movie.MovieGetMovieRes, err error)
	CreateMovie(ctx context.Context, req *movie.MovieCreateMovieReq) (res *movie.MovieCreateMovieRes, err error)
	UpdateMovie(ctx context.Context, req *movie.MovieUpdateMovieReq) (res *movie.MovieUpdateMovieRes, err error)
	DeleteMovie(ctx context.Context, req *movie.MovieDeleteMovieReq) (res *movie.MovieDeleteMovieRes, err error)
	
	// Movie types management
	GetMovieTypes(ctx context.Context, req *movie.MovieGetMovieTypesReq) (res *movie.MovieGetMovieTypesRes, err error)
	CreateMovieType(ctx context.Context, req *movie.MovieCreateMovieTypeReq) (res *movie.MovieCreateMovieTypeRes, err error)
	UpdateMovieType(ctx context.Context, req *movie.MovieUpdateMovieTypeReq) (res *movie.MovieUpdateMovieTypeRes, err error)
	DeleteMovieType(ctx context.Context, req *movie.MovieDeleteMovieTypeReq) (res *movie.MovieDeleteMovieTypeRes, err error)
}
