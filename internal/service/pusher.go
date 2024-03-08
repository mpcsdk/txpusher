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
	IPusher interface {
		GetRecentMsg(ctx context.Context, userId string, addr string) (int, []*entity.Txs)
		// /
		Ack(ctx context.Context, userId string, seqId int) error
		PushTxs(ctx context.Context, chainId string, txs []*entity.Txs)
	}
)

var (
	localPusher IPusher
)

func Pusher() IPusher {
	if localPusher == nil {
		panic("implement not found for interface IPusher, forgot register?")
	}
	return localPusher
}

func RegisterPusher(i IPusher) {
	localPusher = i
}
