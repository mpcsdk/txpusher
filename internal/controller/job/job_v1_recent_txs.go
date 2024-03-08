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

func (c *ControllerV1) RecentTxs(ctx context.Context, req *v1.RecentTxsReq) (*v1.RecentTxsRes, error) {
	info, err := userInfoGeter.Geter(conf.Config.UserTokenUrl).GetUserInfo(ctx, req.Token)
	if err != nil {
		g.Log().Warning(ctx, "get user info failed, err:", err)
		return nil, mpccode.CodeTokenInvalid()
	}
	seq, data := service.Pusher().GetRecentMsg(ctx, info.UserId, "0x1fc35B79FB11Ea7D4532dA128DfA9Db573C51b09")
	res := &v1.RecentTxsRes{
		Seq:  seq,
		Data: data,
	}
	return res, nil
}
