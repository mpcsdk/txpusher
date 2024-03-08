// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	ITxSync interface{}
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
