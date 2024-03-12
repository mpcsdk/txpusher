package db

import (
	"context"
	"time"
	"txpusher/internal/dao"
	"txpusher/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

func (s *sDB) GetContractRule(ctx context.Context, chainId string, address string, flush bool) (*entity.Contractrule, error) {
	rst, err := dao.Contractrule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: func() time.Duration {
			if flush {
				return -1
			} else {
				return s.dbDuration
			}
		}(),
		Name:  dao.Contractrule.Table() + chainId + address,
		Force: true,
	}).Where(dao.Contractrule.Columns().ChainId, chainId).
		Where(dao.Contractrule.Columns().ContractAddress, address).One()
	if err != nil {
		return nil, err
	}
	// /
	rule := &entity.Contractrule{}
	rst.Struct(&rule)
	return rule, nil
}

func (s *sDB) GetContractRuleBriefs(ctx context.Context, chainId string, kind string) ([]*entity.Contractrule, error) {
	model := dao.Contractrule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.Contractrule.Table() + "GetContractRuleBriefs" + chainId + kind,
		Force:    true,
	}).Fields(
		dao.Contractrule.Columns().ChainId,
		dao.Contractrule.Columns().ContractAddress,
		dao.Contractrule.Columns().ContractName,
		dao.Contractrule.Columns().ContractKind,
	)
	if chainId != "" {
		model = model.Where(dao.Contractrule.Columns().ChainId, chainId)
	}
	if kind != "" {
		model = model.Where(dao.Contractrule.Columns().ContractKind, kind)
	}
	rst, err := model.All()
	if err != nil {
		return nil, err
	}
	///
	rule := []*entity.Contractrule{}
	rst.Structs(&rule)
	return rule, nil
}
