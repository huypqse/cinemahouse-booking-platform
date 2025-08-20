// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentDao is the data access object for the table comment.
type CommentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CommentColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CommentColumns defines and stores column names for the table comment.
type CommentColumns struct {
	Id      string //
	Content string //
	Status  string //
	MovieId string //
}

// commentColumns holds the columns for the table comment.
var commentColumns = CommentColumns{
	Id:      "id",
	Content: "content",
	Status:  "status",
	MovieId: "movie_id",
}

// NewCommentDao creates and returns a new DAO object for table data access.
func NewCommentDao(handlers ...gdb.ModelHandler) *CommentDao {
	return &CommentDao{
		group:    "default",
		table:    "comment",
		columns:  commentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CommentDao) Columns() CommentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CommentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
