package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"auth-svc/api/admin/admin"
)

func (c *ControllerAdmin) AdminGetAuditLogs(ctx context.Context, req *admin.AdminGetAuditLogsReq) (res *admin.AdminGetAuditLogsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
