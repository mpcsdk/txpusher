// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"txpusher/internal/dao/internal"
)

// internalChainStatDao is internal type for wrapping internal DAO implements.
type internalChainStatDao = *internal.ChainStatDao

// chainStatDao is the data access object for table chain_stat.
// You can define custom methods on it to extend its functionality as you wish.
type chainStatDao struct {
	internalChainStatDao
}

var (
	// ChainStat is globally public accessible object for table chain_stat operations.
	ChainStat = chainStatDao{
		internal.NewChainStatDao(),
	}
)

// Fill with you ideas below.
