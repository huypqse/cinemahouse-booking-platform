// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomTypeDao is the data access object for the table room_type.
type RoomTypeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  RoomTypeColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// RoomTypeColumns defines and stores column names for the table room_type.
type RoomTypeColumns struct {
	Id     string //
	Name   string //
	Status string //
}

// roomTypeColumns holds the columns for the table room_type.
var roomTypeColumns = RoomTypeColumns{
	Id:     "id",
	Name:   "name",
	Status: "status",
}

// NewRoomTypeDao creates and returns a new DAO object for table data access.
func NewRoomTypeDao(handlers ...gdb.ModelHandler) *RoomTypeDao {
	return &RoomTypeDao{
		group:    "default",
		table:    "room_type",
		columns:  roomTypeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoomTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoomTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoomTypeDao) Columns() RoomTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoomTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoomTypeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RoomTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
