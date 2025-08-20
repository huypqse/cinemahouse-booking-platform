package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/user/user"
)

func (c *ControllerUser) UserLogin(ctx context.Context, req *user.UserLoginReq) (res *user.UserLoginRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
