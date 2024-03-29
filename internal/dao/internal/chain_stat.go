// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChainStatDao is the data access object for table chain_stat.
type ChainStatDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ChainStatColumns // columns contains all the column names of Table for convenient usage.
}

// ChainStatColumns defines and stores column names for table chain_stat.
type ChainStatColumns struct {
	ChainId     string //
	BlockNumber string //
}

// chainStatColumns holds the columns for table chain_stat.
var chainStatColumns = ChainStatColumns{
	ChainId:     "chain_id",
	BlockNumber: "block_number",
}

// NewChainStatDao creates and returns a new DAO object for table data access.
func NewChainStatDao() *ChainStatDao {
	return &ChainStatDao{
		group:   "default",
		table:   "chain_stat",
		columns: chainStatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChainStatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChainStatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChainStatDao) Columns() ChainStatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChainStatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChainStatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChainStatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
