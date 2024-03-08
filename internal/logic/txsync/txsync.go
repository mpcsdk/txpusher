package txsync

import (
	"context"
	"time"
	"txpusher/internal/conf"
	"txpusher/internal/model/pushermod"
	"txpusher/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type sTxSync struct {
	ctx context.Context

	///
	chainStat map[string]int
}

// //
type chainStat struct {
	Status int            `json:"status"`
	Result map[string]int `json:"result"`
}

func (s *sTxSync) getChainsStat() map[string]int {
	v := g.Client().GetVar(s.ctx, conf.Config.Txdatarpc+"/state", nil)

	reps := chainStat{}
	v.Scan(&reps)
	//
	if reps.Status != 0 {
		g.Log().Warning(s.ctx, "syncChainsStat error:", v)
		return nil
	}
	g.Log().Notice(s.ctx, "syncChainsStat ok:", reps.Result)
	return reps.Result
}

// /
type QueryAdvRep struct {
	Status int                     `json:"status"`
	Result []*pushermod.Tansaction `json:"result"`
}

func (s *sTxSync) syncTxs() map[string][]*pushermod.Tansaction {
	txs := map[string][]*pushermod.Tansaction{}

	////get txs
	for chainid, latestNr := range s.chainStat {
		v := g.Client().GetVar(s.ctx, conf.Config.Txdatarpc+"/queryAdv", map[string]interface{}{
			"chainId":   chainid,
			"fromBlock": latestNr + 1,
		})
		resp := QueryAdvRep{}
		v.Scan(&resp)
		if resp.Status != 0 {
			g.Log().Warning(s.ctx, "syncTxs error:", resp, "chainid:", chainid, "fromBlock:", latestNr)
			continue
		}
		//
		///
		if len(resp.Result) == 0 {
			continue
		}
		//
		for _, tx := range resp.Result {
			if tx.Height > latestNr {
				latestNr = tx.Height
			}
		}
		///
		err := service.DB().UpChainStat(s.ctx, chainid, latestNr)
		if err != nil {
			g.Log().Warning(s.ctx, "syncTxs error:", err, "chainid:", chainid, "fromBlock:", latestNr)
			continue
		}
		s.chainStat[chainid] = latestNr
		txs[chainid] = resp.Result
		////
	}
	return txs
}

// /
func (s *sTxSync) run() {
	go func() {
		for {
			data := s.syncTxs()
			for chainid, txs := range data {
				service.Pusher().PushTxs(s.ctx, chainid, pushermod.Trans2Txs(txs))
				///
			}
			time.Sleep(time.Second * 10)
		}
	}()
}

func new() *sTxSync {
	s := &sTxSync{
		ctx:       gctx.GetInitCtx(),
		chainStat: map[string]int{},
	}

	//
	chainstat, err := service.DB().GetChainStat(s.ctx)
	if err != nil {
		panic(err)
	}
	for _, c := range chainstat {
		s.chainStat[c.ChainId] = c.BlockNumber
	}
	///sync chainstat
	remoteChainStat := s.getChainsStat()
	for chainid, latestNr := range remoteChainStat {
		if _, ok := s.chainStat[chainid]; !ok {
			s.chainStat[chainid] = latestNr
		}
	}
	s.run()
	return s
}

func init() {
	service.RegisterTxSync(new())
}
