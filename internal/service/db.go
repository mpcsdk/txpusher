// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"txpusher/internal/model/entity"
)

type (
	IDB interface {
		GetContractRule(ctx context.Context, chainId string, address string, flush bool) (*entity.Contractrule, error)
		GetContractRuleBriefs(ctx context.Context, chainId string, kind string) ([]*entity.Contractrule, error)
		GetOfflineMsg(ctx context.Context, userId string) ([]*entity.Txs, error)
		InsertOfflineMsg(ctx context.Context, userId string, tx *entity.Txs) error
		DeleteOfflineMsg(ctx context.Context, userId string, chainId string, latestNr int) error
		GetChainStat(ctx context.Context) ([]*entity.ChainStat, error)
		UpChainStat(ctx context.Context, chainId string, latestNr int) error
		// /
		InsertTxs(ctx context.Context, txs []*entity.Txs) error
		GetTxs(ctx context.Context, chainId string, blockNumber int, userId string) ([]*entity.Txs, error)
	}
)

var (
	localDB IDB
)

func DB() IDB {
	if localDB == nil {
		panic("implement not found for interface IDB, forgot register?")
	}
	return localDB
}

func RegisterDB(i IDB) {
	localDB = i
}
