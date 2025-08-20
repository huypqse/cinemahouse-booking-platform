// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SeatDao is the data access object for the table seat.
type SeatDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SeatColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SeatColumns defines and stores column names for the table seat.
type SeatColumns struct {
	Id              string //
	SeatRow         string //
	SeatColumn      string //
	Status          string //
	ScreeningRoomId string //
	SeatTypeId      string //
}

// seatColumns holds the columns for the table seat.
var seatColumns = SeatColumns{
	Id:              "id",
	SeatRow:         "seat_row",
	SeatColumn:      "seat_column",
	Status:          "status",
	ScreeningRoomId: "screening_room_id",
	SeatTypeId:      "seat_type_id",
}

// NewSeatDao creates and returns a new DAO object for table data access.
func NewSeatDao(handlers ...gdb.ModelHandler) *SeatDao {
	return &SeatDao{
		group:    "default",
		table:    "seat",
		columns:  seatColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SeatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SeatDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SeatDao) Columns() SeatColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SeatDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SeatDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SeatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
