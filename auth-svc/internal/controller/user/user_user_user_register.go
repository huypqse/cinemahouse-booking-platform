package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/user/user"
)

func (c *ControllerUser) UserRegister(ctx context.Context, req *user.UserRegisterReq) (res *user.UserRegisterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
