package cinema

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"cinemahouse-svc/api/cinema/cinema"
)

func (c *ControllerCinema) CinemaDeleteRoomType(ctx context.Context, req *cinema.CinemaDeleteRoomTypeReq) (res *cinema.CinemaDeleteRoomTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
