// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/mpcsdk/mpcCommon/mq"
)

type (
	ITxSync interface {
		NotifyContractRule(ctx context.Context, req *mq.ContractRuleReq) error
	}
)

var (
	localTxSync ITxSync
)

func TxSync() ITxSync {
	if localTxSync == nil {
		panic("implement not found for interface ITxSync, forgot register?")
	}
	return localTxSync
}

func RegisterTxSync(i ITxSync) {
	localTxSync = i
}
