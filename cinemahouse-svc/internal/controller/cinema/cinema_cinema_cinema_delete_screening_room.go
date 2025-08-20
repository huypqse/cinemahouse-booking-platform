package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaDeleteScreeningRoom(ctx context.Context, req *cinema.CinemaDeleteScreeningRoomReq) (res *cinema.CinemaDeleteScreeningRoomRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
