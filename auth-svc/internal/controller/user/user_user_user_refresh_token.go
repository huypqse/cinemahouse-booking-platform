package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/user/user"
)

func (c *ControllerUser) UserRefreshToken(ctx context.Context, req *user.UserRefreshTokenReq) (res *user.UserRefreshTokenRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
