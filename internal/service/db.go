// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IDB interface {
		FetchJobStatus(ctx context.Context, userId string) (int, error)
		UpJobStatus(ctx context.Context, userId string, seqId string) error
		FetchTxsBySeqId(ctx context.Context, seqId string) []string
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
