package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/admin/admin"
)

func (c *ControllerAdmin) AdminUnbanUser(ctx context.Context, req *admin.AdminUnbanUserReq) (res *admin.AdminUnbanUserRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
