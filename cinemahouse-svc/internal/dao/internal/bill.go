// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BillDao is the data access object for the table bill.
type BillDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  BillColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// BillColumns defines and stores column names for the table bill.
type BillColumns struct {
	Id          string //
	BillDate    string //
	BillTime    string //
	TotalAmount string //
	Status      string //
	UserId      string //
	PaymentId   string //
	RequestId   string //
}

// billColumns holds the columns for the table bill.
var billColumns = BillColumns{
	Id:          "id",
	BillDate:    "bill_date",
	BillTime:    "bill_time",
	TotalAmount: "total_amount",
	Status:      "status",
	UserId:      "user_id",
	PaymentId:   "payment_id",
	RequestId:   "request_id",
}

// NewBillDao creates and returns a new DAO object for table data access.
func NewBillDao(handlers ...gdb.ModelHandler) *BillDao {
	return &BillDao{
		group:    "default",
		table:    "bill",
		columns:  billColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BillDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BillDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BillDao) Columns() BillColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BillDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BillDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BillDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
