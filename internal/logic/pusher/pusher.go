package pusher

import (
	"context"
	"time"
	"txpusher/internal/conf"
	"txpusher/internal/model/entity"
	"txpusher/internal/service"

	"github.com/gogf/gf/v2/os/gctx"
)

type sPusher struct {
	ctx context.Context
	///
	jobs       map[string]*Jober
	addrUserId map[string]string
	chainStat  map[string]uint64
	///
	jobExpire time.Duration
}

func (s *sPusher) GetRecentMsg(ctx context.Context, userId string, addr string) (int, []*entity.Txs) {
	///	///
	if j, ok := s.jobs[userId]; ok {
		return j.RecentMsg()
	}
	///

	s.jobs[userId] = newJober(s.ctx, userId, addr, s.jobExpire)
	return s.jobs[userId].RecentMsg()
}

// /
func (s *sPusher) Ack(ctx context.Context, userId string, seqId int) error {
	if j, ok := s.jobs[userId]; ok {
		j.Ack(ctx, seqId)
	}
	return nil
}

func (s *sPusher) PushTxs(ctx context.Context, chainId string, txs []*entity.Txs) {
	for _, tx := range txs {
		fromUserId := tx.FromAddr
		toUserId := tx.ToAddr
		//
		if j, ok := s.jobs[fromUserId]; ok {
			j.PushMsg(tx)
		} else {
			service.DB().InsertOfflineMsg(s.ctx, fromUserId, tx)
		}
		if j, ok := s.jobs[toUserId]; ok {
			j.PushMsg(tx)
		} else {
			service.DB().InsertOfflineMsg(s.ctx, toUserId, tx)
		}
	}
}

// /
// sync chains txs
func (s *sPusher) start() {
	go func() {

	}()
}

func (s *sPusher) maintJobs() {
	go func() {
		for {
			for k, j := range s.jobs {
				//判断jober超过有效期
				if !j.IsExpire() {
					s.jobs[k] = nil
					s.addrUserId[j.addr] = ""
				}
			}
			time.Sleep(time.Minute * 1)
		}
	}()
}

// /
func new() *sPusher {
	j := &sPusher{
		ctx:        gctx.GetInitCtx(),
		jobs:       map[string]*Jober{},
		addrUserId: map[string]string{},
		chainStat:  map[string]uint64{},
		jobExpire:  time.Duration(conf.Config.Cache.DBCacheDuration),
	}
	j.start()
	j.maintJobs()
	return j
}
func init() {
	service.RegisterPusher(new())
}
