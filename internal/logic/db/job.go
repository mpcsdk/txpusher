package db

import (
	"context"
	"txpusher/internal/dao"
	"txpusher/internal/model/do"
	"txpusher/internal/model/entity"
)

func (s *sDB) GetOfflineMsg(ctx context.Context, userId string) ([]*entity.Txs, error) {

	rst, err := dao.JobOfflineTxs.Ctx(ctx).Where(do.JobOfflineTxs{
		UserId: userId,
	}).All()
	if err != nil {
		return nil, err
	}
	data := []*entity.Txs{}
	rst.Structs(&data)
	return data, nil
}
func (s *sDB) InsertOfflineMsg(ctx context.Context, userId string, tx *entity.Txs) error {

	_, err := dao.JobOfflineTxs.Ctx(ctx).Data(tx).Insert()
	return err
}
func (s *sDB) DeleteOfflineMsg(ctx context.Context, userId string, chainId string, latestNr int) error {

	_, err := dao.JobOfflineTxs.Ctx(ctx).Where(do.JobOfflineTxs{
		UserId:  userId,
		ChainId: chainId,
	}).WhereLTE(dao.JobOfflineTxs.Columns().Height, latestNr).Delete()

	return err

}
