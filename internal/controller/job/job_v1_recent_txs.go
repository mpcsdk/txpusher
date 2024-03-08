package job

import (
	"context"
	v1 "txpusher/api/job/v1"
	"txpusher/internal/service"
)

func (c *ControllerV1) RecentTxs(ctx context.Context, req *v1.RecentTxsReq) (*v1.RecentTxsRes, error) {
	seq, data := service.Pusher().GetRecentMsg(ctx, "userid", "0x1fc35B79FB11Ea7D4532dA128DfA9Db573C51b09")
	res := &v1.RecentTxsRes{
		Seq:  seq,
		Data: data,
	}
	return res, nil
}
