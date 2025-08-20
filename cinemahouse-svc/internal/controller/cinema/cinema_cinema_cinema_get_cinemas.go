package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaGetCinemas(ctx context.Context, req *cinema.CinemaGetCinemasReq) (res *cinema.CinemaGetCinemasRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
