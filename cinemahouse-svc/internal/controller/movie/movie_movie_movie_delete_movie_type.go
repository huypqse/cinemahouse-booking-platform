package movie

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/movie/movie"
)

func (c *ControllerMovie) MovieDeleteMovieType(ctx context.Context, req *movie.MovieDeleteMovieTypeReq) (res *movie.MovieDeleteMovieTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
