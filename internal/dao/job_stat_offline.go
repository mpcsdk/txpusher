// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"txpusher/internal/dao/internal"
)

// internalJobStatOfflineDao is internal type for wrapping internal DAO implements.
type internalJobStatOfflineDao = *internal.JobStatOfflineDao

// jobStatOfflineDao is the data access object for table job_stat_offline.
// You can define custom methods on it to extend its functionality as you wish.
type jobStatOfflineDao struct {
	internalJobStatOfflineDao
}

var (
	// JobStatOffline is globally public accessible object for table job_stat_offline operations.
	JobStatOffline = jobStatOfflineDao{
		internal.NewJobStatOfflineDao(),
	}
)

// Fill with you ideas below.
