// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package job

import (
	"context"
	
	"txpusher/api/job/v1"
)

type IJobV1 interface {
	RecentTxs(ctx context.Context, req *v1.RecentTxsReq) (res *v1.RecentTxsRes, err error)
	TxsAck(ctx context.Context, req *v1.TxsAckReq) (res *v1.TxsAckRes, err error)
}


