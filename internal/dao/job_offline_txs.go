// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"txpusher/internal/dao/internal"
)

// internalJobOfflineTxsDao is internal type for wrapping internal DAO implements.
type internalJobOfflineTxsDao = *internal.JobOfflineTxsDao

// jobOfflineTxsDao is the data access object for table job_offline_txs.
// You can define custom methods on it to extend its functionality as you wish.
type jobOfflineTxsDao struct {
	internalJobOfflineTxsDao
}

var (
	// JobOfflineTxs is globally public accessible object for table job_offline_txs operations.
	JobOfflineTxs = jobOfflineTxsDao{
		internal.NewJobOfflineTxsDao(),
	}
)

// Fill with you ideas below.
