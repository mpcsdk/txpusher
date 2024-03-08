package db

import (
	"context"
	"txpusher/internal/dao"
	"txpusher/internal/model/do"
	"txpusher/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sDB) GetChainStat(ctx context.Context) ([]*entity.ChainStat, error) {
	rst, err := dao.ChainStat.Ctx(ctx).All()
	if err != nil {
		g.Log().Warning(ctx, "FetchChainStat failed, err:", err)
		return nil, err
	}
	///
	stat := []*entity.ChainStat{}
	rst.Structs(&stat)
	return stat, nil
}
func (s *sDB) UpChainStat(ctx context.Context, chainId string, latestNr int) error {

	_, err := dao.ChainStat.Ctx(ctx).
		Data(do.ChainStat{
			ChainId:     chainId,
			BlockNumber: latestNr,
		}).
		OnConflict(dao.ChainStat.Columns().ChainId).Save()
	return err
}

// /
func (s *sDB) InsertTxs(ctx context.Context, txs []*entity.Txs) error {
	_, err := dao.Txs.Ctx(ctx).Data(txs).Batch(10).Insert()
	return err
}

func (s *sDB) GetTxs(ctx context.Context, chainId string, blockNumber int, userId string) ([]*entity.Txs, error) {
	return nil, nil
}
