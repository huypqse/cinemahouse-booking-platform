package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/auth/auth"
)

func (c *ControllerAuth) AuthEnableTwoFactor(ctx context.Context, req *auth.AuthEnableTwoFactorReq) (res *auth.AuthEnableTwoFactorRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
