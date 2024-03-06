// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package job

import (
	"context"
	
	"txpusher/api/job/v1"
)

type IJobV1 interface {
	FetchTxs(ctx context.Context, req *v1.FetchTxsReq) (res *v1.FetchTxsRes, err error)
	FetchAck(ctx context.Context, req *v1.FetchAckReq) (res *v1.FetchAckRes, err error)
}


