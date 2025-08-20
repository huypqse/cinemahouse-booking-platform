package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/auth/auth"
)

func (c *ControllerAuth) AuthFacebookAuth(ctx context.Context, req *auth.AuthFacebookAuthReq) (res *auth.AuthFacebookAuthRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
