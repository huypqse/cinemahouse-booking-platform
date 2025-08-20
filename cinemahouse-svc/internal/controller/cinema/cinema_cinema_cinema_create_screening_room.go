package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaCreateScreeningRoom(ctx context.Context, req *cinema.CinemaCreateScreeningRoomReq) (res *cinema.CinemaCreateScreeningRoomRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
