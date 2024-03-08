// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// JobOfflineTxsDao is the data access object for table job_offline_txs.
type JobOfflineTxsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns JobOfflineTxsColumns // columns contains all the column names of Table for convenient usage.
}

// JobOfflineTxsColumns defines and stores column names for table job_offline_txs.
type JobOfflineTxsColumns struct {
	ChainId   string //
	Height    string //
	BlockHash string //
	Ts        string //
	TxHash    string //
	FromAddr  string //
	ToAddr    string //
	Value     string //
	Contract  string //
	Gas       string //
	GasPrice  string //
	TxIdx     string //
	LogIdx    string //
	UserId    string //
}

// jobOfflineTxsColumns holds the columns for table job_offline_txs.
var jobOfflineTxsColumns = JobOfflineTxsColumns{
	ChainId:   "chain_id",
	Height:    "height",
	BlockHash: "block_hash",
	Ts:        "ts",
	TxHash:    "tx_hash",
	FromAddr:  "from_addr",
	ToAddr:    "to_addr",
	Value:     "value",
	Contract:  "contract",
	Gas:       "gas",
	GasPrice:  "gas_price",
	TxIdx:     "tx_idx",
	LogIdx:    "log_idx",
	UserId:    "user_id",
}

// NewJobOfflineTxsDao creates and returns a new DAO object for table data access.
func NewJobOfflineTxsDao() *JobOfflineTxsDao {
	return &JobOfflineTxsDao{
		group:   "default",
		table:   "job_offline_txs",
		columns: jobOfflineTxsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *JobOfflineTxsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *JobOfflineTxsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *JobOfflineTxsDao) Columns() JobOfflineTxsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *JobOfflineTxsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *JobOfflineTxsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *JobOfflineTxsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
