// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package movie

import (
	"context"

	"cinemahouse-svc/api/movie/movie"
)

type IMovieMovie interface {
	MovieGetMovies(ctx context.Context, req *movie.MovieGetMoviesReq) (res *movie.MovieGetMoviesRes, err error)
	MovieGetMovie(ctx context.Context, req *movie.MovieGetMovieReq) (res *movie.MovieGetMovieRes, err error)
	MovieCreateMovie(ctx context.Context, req *movie.MovieCreateMovieReq) (res *movie.MovieCreateMovieRes, err error)
	MovieUpdateMovie(ctx context.Context, req *movie.MovieUpdateMovieReq) (res *movie.MovieUpdateMovieRes, err error)
	MovieDeleteMovie(ctx context.Context, req *movie.MovieDeleteMovieReq) (res *movie.MovieDeleteMovieRes, err error)
	MovieGetMovieTypes(ctx context.Context, req *movie.MovieGetMovieTypesReq) (res *movie.MovieGetMovieTypesRes, err error)
	MovieCreateMovieType(ctx context.Context, req *movie.MovieCreateMovieTypeReq) (res *movie.MovieCreateMovieTypeRes, err error)
	MovieUpdateMovieType(ctx context.Context, req *movie.MovieUpdateMovieTypeReq) (res *movie.MovieUpdateMovieTypeRes, err error)
	MovieDeleteMovieType(ctx context.Context, req *movie.MovieDeleteMovieTypeReq) (res *movie.MovieDeleteMovieTypeRes, err error)
}
