package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/user/user"
)

func (c *ControllerUser) UserProfile(ctx context.Context, req *user.UserProfileReq) (res *user.UserProfileRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
