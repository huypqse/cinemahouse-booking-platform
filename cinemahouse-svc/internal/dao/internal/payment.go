// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PaymentDao is the data access object for the table payment.
type PaymentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PaymentColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PaymentColumns defines and stores column names for the table payment.
type PaymentColumns struct {
	Id            string //
	PaymentMethod string //
	PaymentStatus string //
	PaymentDate   string //
	Amount        string //
	TransactionId string //
}

// paymentColumns holds the columns for the table payment.
var paymentColumns = PaymentColumns{
	Id:            "id",
	PaymentMethod: "payment_method",
	PaymentStatus: "payment_status",
	PaymentDate:   "payment_date",
	Amount:        "amount",
	TransactionId: "transaction_id",
}

// NewPaymentDao creates and returns a new DAO object for table data access.
func NewPaymentDao(handlers ...gdb.ModelHandler) *PaymentDao {
	return &PaymentDao{
		group:    "default",
		table:    "payment",
		columns:  paymentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PaymentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PaymentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PaymentDao) Columns() PaymentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PaymentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PaymentDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PaymentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
