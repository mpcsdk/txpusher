package job

import (
	"context"

	v1 "txpusher/api/job/v1"
	"txpusher/internal/service"
)

func (c *ControllerV1) TxsAck(ctx context.Context, req *v1.TxsAckReq) (res *v1.TxsAckRes, err error) {
	service.Pusher().Ack(ctx, "userid", req.Seq)
	return nil, nil
}
