package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaCreateCinema(ctx context.Context, req *cinema.CinemaCreateCinemaReq) (res *cinema.CinemaCreateCinemaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
