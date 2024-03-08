package pusher

import (
	"context"
	"sync"
	"time"
	"txpusher/internal/model/entity"
	"txpusher/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type Jober struct {
	///
	lock sync.Mutex
	///
	ctx    context.Context
	userId string
	addr   string
	//
	msg []*entity.Txs
	//
	seqId             int
	oldmsg            []*entity.Txs
	oldmsgchainlatest map[string]int
	///
	expire time.Time
	//
}

func (s *Jober) PushMsg(msg *entity.Txs) {
	s.lock.Lock()
	s.msg = append(s.msg, msg)
	s.lock.Unlock()
}

// /
func (s *Jober) RecentMsg() (int, []*entity.Txs) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.msg) == 0 {
		return s.seqId, nil
	}
	///
	if len(s.oldmsg) == 0 {
		if s.seqId > 1000000 {
			s.seqId = 0
		}
		s.seqId += 1
		s.oldmsg = s.msg
		s.oldmsgchainlatest = map[string]int{}
		for latest, tx := range s.oldmsg {
			if _, ok := s.oldmsgchainlatest[tx.ChainId]; !ok {
				s.oldmsgchainlatest[tx.ChainId] = tx.Height
			} else {
				if latest < tx.Height {
					latest = tx.Height
					s.oldmsgchainlatest[tx.ChainId] = latest
				}
			}
		}
		s.msg = nil
	}
	///
	return s.seqId, s.oldmsg
}

func (s *Jober) Ack(ctx context.Context, seqId int) {
	if s.seqId == seqId {
		s.lock.Lock()
		defer s.lock.Unlock()
		//

		for chainid, latestNr := range s.oldmsgchainlatest {
			err := service.DB().DeleteOfflineMsg(ctx, s.userId, chainid, latestNr)
			if err != nil {
				g.Log().Warning(s.ctx, "DeleteOfflineMsg err", err)
			}
		}
		//
		s.oldmsg = nil
		s.oldmsgchainlatest = nil
	}
}
func (s *Jober) IsExpire() bool {
	return time.Now().After(s.expire)
}

// /
func newJober(ctx context.Context, userId, addr string, expire time.Duration) *Jober {
	offlineMsg, err := service.DB().GetOfflineMsg(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "newJober err", err)
	}
	////
	s := &Jober{
		ctx:    ctx,
		userId: userId,
		addr:   addr,
		////
		msg:    []*entity.Txs{},
		oldmsg: []*entity.Txs{},
		seqId:  0,
		///
		expire: time.Now().Add(expire),
	}
	//offline txs
	s.msg = offlineMsg

	return s
}
