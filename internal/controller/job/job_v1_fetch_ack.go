package job

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"txpusher/api/job/v1"
)

func (c *ControllerV1) FetchAck(ctx context.Context, req *v1.FetchAckReq) (res *v1.FetchAckRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
