package job

import (
	"context"
	"time"
	"txpusher/internal/conf"

	"github.com/gogf/gf/v2/frame/g"
)

type sJob struct {
	ctx context.Context
	///
	jobs       map[string]*Jober
	addrUserId map[string]string
	chainStat  map[uint64]uint64
	///
	jobExpire time.Duration
}

func (s *sJob) FetchMsg(ctx context.Context, userId string) (int, []*item) {
	if j, ok := s.jobs[userId]; ok {
		return j.RecentMsg()
	}
	return 0, nil
}

// /
func (s *sJob) Ack(ctx context.Context, userId string, seqId int) error {
	if j, ok := s.jobs[userId]; ok {
		j.Ack(seqId)
	}
	return nil
}

// //
type chainStat struct {
}

func (s *sJob) syncChainsStat() {
	v := g.Client().GetVar(s.ctx, conf.Config.Txdatarpc+"/chainState", nil)

	reps := chainStat{}
	v.Scan(&reps)
	//
}

// /
type QueryAdvRep struct {
	Status int     `json:"status"`
	Result []*item `json:"result"`
}
type item struct {
	ChainId   uint64 `json:"chainId"`
	Height    string `json:"height"`
	Blockhash string `json:"blockhash"`
	Ts        string `json:"timestamp"`
	Txhash    string `json:"txhash"`
	Toaddr    string `json:"toaddr"`
	Fromaddr  string `json:"fromaddr"`
	Value     string `json:"value"`
	Contract  string `json:"contract"`
	Gas       string `json:"gas"`
	Gasprice  string `json:"gasprice"`
}

func (s *sJob) syncTxs() []*item {
	data := []*item{}
	for chainid, blockNr := range s.chainStat {

		v := g.Client().GetVar(s.ctx, conf.Config.Txdatarpc, map[string]interface{}{
			"chainid":     chainid,
			"startNumber": blockNr,
		})
		reps := QueryAdvRep{}
		v.Scan(&reps)
		data = append(data, reps.Result...)
	}
	return data
}

// /
// sync chains txs
func (s *sJob) start() {
	go func() {
		for {
			txs := s.syncTxs()
			for _, tx := range txs {
				fromUserId := tx.Fromaddr
				toUserId := tx.Toaddr
				//
				if j, ok := s.jobs[fromUserId]; ok {

				} else {
					j = newJober(s.jobExpire)
					s.jobs[fromUserId] = j
				}
				if j, ok := s.jobs[toUserId]; ok {
				} else {
					j = newJober(s.jobExpire)
					s.jobs[toUserId] = j
				}
				//
				s.jobs[fromUserId].PushMsg(nil)
				s.jobs[toUserId].PushMsg(nil)
			}
		}
	}()
}

func (s *sJob) maintJobs() {
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
func new() *sJob {
	j := &sJob{}
	j.start()
	j.maintJobs()
	return j
}
func init() {

}
