package movie

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/movie/movie"
)

func (c *ControllerMovie) MovieUpdateMovie(ctx context.Context, req *movie.MovieUpdateMovieReq) (res *movie.MovieUpdateMovieRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
