package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/admin/admin"
)

func (c *ControllerAdmin) AdminGetPermissions(ctx context.Context, req *admin.AdminGetPermissionsReq) (res *admin.AdminGetPermissionsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
