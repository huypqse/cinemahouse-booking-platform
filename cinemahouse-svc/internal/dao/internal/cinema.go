// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CinemaDao is the data access object for the table cinema.
type CinemaDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CinemaColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CinemaColumns defines and stores column names for the table cinema.
type CinemaColumns struct {
	Id            string //
	CinemaName    string //
	CinemaAddress string //
	CinemaStatus  string //
}

// cinemaColumns holds the columns for the table cinema.
var cinemaColumns = CinemaColumns{
	Id:            "id",
	CinemaName:    "cinema_name",
	CinemaAddress: "cinema_address",
	CinemaStatus:  "cinema_status",
}

// NewCinemaDao creates and returns a new DAO object for table data access.
func NewCinemaDao(handlers ...gdb.ModelHandler) *CinemaDao {
	return &CinemaDao{
		group:    "default",
		table:    "cinema",
		columns:  cinemaColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CinemaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CinemaDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CinemaDao) Columns() CinemaColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CinemaDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CinemaDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CinemaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
