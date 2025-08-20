package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

// System Management APIs
type AdminGetSystemStatsReq struct {
	g.Meta `path:"/admin/system/stats" tags:"Admin" method:"get" summary:"Get system statistics"`
}

type AdminGetSystemStatsRes struct {
	TotalUsers           int64 `json:"total_users"`
	ActiveUsers          int64 `json:"active_users"`
	InactiveUsers        int64 `json:"inactive_users"`
	BannedUsers          int64 `json:"banned_users"`
	TotalRoles           int64 `json:"total_roles"`
	TotalPermissions     int64 `json:"total_permissions"`
	InvalidatedTokens    int64 `json:"invalidated_tokens"`
	ExpiredTokens        int64 `json:"expired_tokens"`
	UsersRegisteredToday int64 `json:"users_registered_today"`
	UsersLoginToday      int64 `json:"users_login_today"`
}

type AdminGetInvalidatedTokensReq struct {
	g.Meta   `path:"/admin/system/invalidated-tokens" tags:"Admin" method:"get" summary:"Get invalidated tokens"`
	Page     int `json:"page" v:"min:1#Page must be at least 1"`
	PageSize int `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
}

type AdminGetInvalidatedTokensRes struct {
	Tokens     []InvalidatedTokenInfo `json:"tokens"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
	TotalPages int                    `json:"total_pages"`
}

type InvalidatedTokenInfo struct {
	TokenId    string `json:"token_id"`
	ExpiryTime string `json:"expiry_time"`
	CreatedAt  string `json:"created_at"`
}

type AdminCleanupExpiredTokensReq struct {
	g.Meta `path:"/admin/system/cleanup-tokens" tags:"Admin" method:"post" summary:"Cleanup expired tokens"`
}

type AdminCleanupExpiredTokensRes struct {
	DeletedCount int64  `json:"deleted_count"`
	Message      string `json:"message"`
}

// Audit log APIs
type AdminGetAuditLogsReq struct {
	g.Meta    `path:"/admin/system/audit-logs" tags:"Admin" method:"get" summary:"Get audit logs"`
	Page      int    `json:"page" v:"min:1#Page must be at least 1"`
	PageSize  int    `json:"page_size" v:"min:1|max:100#Page size must be between 1 and 100"`
	UserId    int64  `json:"user_id"`
	Action    string `json:"action"`
	StartDate string `json:"start_date" v:"date#Invalid start date format"`
	EndDate   string `json:"end_date" v:"date#Invalid end date format"`
}

type AdminGetAuditLogsRes struct {
	Logs       []AuditLogInfo `json:"logs"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int            `json:"total_pages"`
}

type AuditLogInfo struct {
	LogId     int64  `json:"log_id"`
	UserId    int64  `json:"user_id"`
	Username  string `json:"username"`
	Action    string `json:"action"`
	Resource  string `json:"resource"`
	Details   string `json:"details"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	CreatedAt string `json:"created_at"`
}
