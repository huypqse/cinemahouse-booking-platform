// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Bill is the golang structure for table bill.
type Bill struct {
	Id          int64       `json:"id"          orm:"id"           description:""`
	BillDate    *gtime.Time `json:"billDate"    orm:"bill_date"    description:""`
	BillTime    *gtime.Time `json:"billTime"    orm:"bill_time"    description:""`
	TotalAmount float64     `json:"totalAmount" orm:"total_amount" description:""`
	Status      string      `json:"status"      orm:"status"       description:""`
	UserId      int64       `json:"userId"      orm:"user_id"      description:""`
	PaymentId   int64       `json:"paymentId"   orm:"payment_id"   description:""`
	RequestId   string      `json:"requestId"   orm:"request_id"   description:""`
}
