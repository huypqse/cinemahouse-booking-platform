// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Payment is the golang structure for table payment.
type Payment struct {
	Id            int64       `json:"id"            orm:"id"             description:""`
	PaymentMethod string      `json:"paymentMethod" orm:"payment_method" description:""`
	PaymentStatus string      `json:"paymentStatus" orm:"payment_status" description:""`
	PaymentDate   *gtime.Time `json:"paymentDate"   orm:"payment_date"   description:""`
	Amount        float64     `json:"amount"        orm:"amount"         description:""`
	TransactionId string      `json:"transactionId" orm:"transaction_id" description:""`
}
