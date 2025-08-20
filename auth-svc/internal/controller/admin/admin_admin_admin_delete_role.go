package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/admin/admin"
)

func (c *ControllerAdmin) AdminDeleteRole(ctx context.Context, req *admin.AdminDeleteRoleReq) (res *admin.AdminDeleteRoleRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
