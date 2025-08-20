package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaDeleteCinema(ctx context.Context, req *cinema.CinemaDeleteCinemaReq) (res *cinema.CinemaDeleteCinemaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
