package txsync

import (
	"context"
	"txpusher/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mq"
)

func (s *sTxSync) NotifyContractRule(ctx context.Context, req *mq.ContractRuleReq) error {
	brief, err := service.DB().GetContractRuleBriefs(ctx, req.ChainId, req.ContractAddress)
	if err != nil {
		g.Log().Error(ctx, "get contract rule err", err)
		return err
	}
	///
	if brief != nil {
		s.contractChain[req.ContractAddress] = req.ChainId
	}
	///
	return nil
}
