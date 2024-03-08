// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TxsOfflineDao is the data access object for table txs_offline.
type TxsOfflineDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TxsOfflineColumns // columns contains all the column names of Table for convenient usage.
}

// TxsOfflineColumns defines and stores column names for table txs_offline.
type TxsOfflineColumns struct {
	UserId      string //
	ChainId     string //
	BlockNumber string //
}

// txsOfflineColumns holds the columns for table txs_offline.
var txsOfflineColumns = TxsOfflineColumns{
	UserId:      "user_id",
	ChainId:     "chain_id",
	BlockNumber: "block_number",
}

// NewTxsOfflineDao creates and returns a new DAO object for table data access.
func NewTxsOfflineDao() *TxsOfflineDao {
	return &TxsOfflineDao{
		group:   "default",
		table:   "txs_offline",
		columns: txsOfflineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TxsOfflineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TxsOfflineDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TxsOfflineDao) Columns() TxsOfflineColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TxsOfflineDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TxsOfflineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TxsOfflineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
