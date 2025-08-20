package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaCreateRoomType(ctx context.Context, req *cinema.CinemaCreateRoomTypeReq) (res *cinema.CinemaCreateRoomTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
