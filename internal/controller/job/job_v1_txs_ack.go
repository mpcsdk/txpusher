package job

import (
	"context"

	v1 "txpusher/api/job/v1"
	"txpusher/internal/conf"
	"txpusher/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/userInfoGeter"
)

func (c *ControllerV1) TxsAck(ctx context.Context, req *v1.TxsAckReq) (res *v1.TxsAckRes, err error) {
	info, err := userInfoGeter.Geter(conf.Config.UserTokenUrl).GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "get user info failed, err:", err)
		return nil, mpccode.CodeTokenInvalid()
	}
	service.Pusher().Ack(ctx, info.UserId, req.Seq)
	return nil, nil
}
