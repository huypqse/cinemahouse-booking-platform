// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MovieDao is the data access object for the table movie.
type MovieDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MovieColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MovieColumns defines and stores column names for the table movie.
type MovieColumns struct {
	Id          string //
	MovieName   string //
	Content     string //
	Duration    string //
	Language    string //
	Director    string //
	Actors      string //
	AgeLimit    string //
	CoverPhoto  string //
	ReleaseDate string //
	StartDate   string //
	Status      string //
}

// movieColumns holds the columns for the table movie.
var movieColumns = MovieColumns{
	Id:          "id",
	MovieName:   "movie_name",
	Content:     "content",
	Duration:    "duration",
	Language:    "language",
	Director:    "director",
	Actors:      "actors",
	AgeLimit:    "age_limit",
	CoverPhoto:  "cover_photo",
	ReleaseDate: "release_date",
	StartDate:   "start_date",
	Status:      "status",
}

// NewMovieDao creates and returns a new DAO object for table data access.
func NewMovieDao(handlers ...gdb.ModelHandler) *MovieDao {
	return &MovieDao{
		group:    "default",
		table:    "movie",
		columns:  movieColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MovieDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MovieDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MovieDao) Columns() MovieColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MovieDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MovieDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MovieDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
