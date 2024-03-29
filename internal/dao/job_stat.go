// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"txpusher/internal/dao/internal"
)

// internalJobStatDao is internal type for wrapping internal DAO implements.
type internalJobStatDao = *internal.JobStatDao

// jobStatDao is the data access object for table job_stat.
// You can define custom methods on it to extend its functionality as you wish.
type jobStatDao struct {
	internalJobStatDao
}

var (
	// JobStat is globally public accessible object for table job_stat operations.
	JobStat = jobStatDao{
		internal.NewJobStatDao(),
	}
)

// Fill with you ideas below.
