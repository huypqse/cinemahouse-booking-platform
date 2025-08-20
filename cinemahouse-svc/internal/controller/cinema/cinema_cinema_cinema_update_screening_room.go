package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaUpdateScreeningRoom(ctx context.Context, req *cinema.CinemaUpdateScreeningRoomReq) (res *cinema.CinemaUpdateScreeningRoomRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
