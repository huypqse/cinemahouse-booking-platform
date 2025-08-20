// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TicketDao is the data access object for the table ticket.
type TicketDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TicketColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TicketColumns defines and stores column names for the table ticket.
type TicketColumns struct {
	Id                 string //
	Status             string //
	ScreeningSessionId string //
	TicketPriceId      string //
	BillId             string //
	SeatId             string //
}

// ticketColumns holds the columns for the table ticket.
var ticketColumns = TicketColumns{
	Id:                 "id",
	Status:             "status",
	ScreeningSessionId: "screening_session_id",
	TicketPriceId:      "ticket_price_id",
	BillId:             "bill_id",
	SeatId:             "seat_id",
}

// NewTicketDao creates and returns a new DAO object for table data access.
func NewTicketDao(handlers ...gdb.ModelHandler) *TicketDao {
	return &TicketDao{
		group:    "default",
		table:    "ticket",
		columns:  ticketColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TicketDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TicketDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TicketDao) Columns() TicketColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TicketDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TicketDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TicketDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
