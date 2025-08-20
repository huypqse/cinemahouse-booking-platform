// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Bill is the golang structure of table bill for DAO operations like Where/Data.
type Bill struct {
	g.Meta      `orm:"table:bill, do:true"`
	Id          interface{} //
	BillDate    *gtime.Time //
	BillTime    *gtime.Time //
	TotalAmount interface{} //
	Status      interface{} //
	UserId      interface{} //
	PaymentId   interface{} //
	RequestId   interface{} //
}
