package v1

import "github.com/gogf/gf/v2/frame/g"

type FetchTxsReq struct {
	g.Meta `path:"/FetchTxs" tags:"FetchTxs" method:"get" summary:"You first hello api"`
	Token  string `json:"token"`
}
type FetchTxsRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Txs    []string `json:"txs"`
}

// //
type FetchAckReq struct {
	g.Meta `path:"/FetchAck" tags:"FetchAck" method:"get" summary:"You first hello api"`
	Token  string `json:"token"`
	Seq    string `json:"seq"`
}
type FetchAckRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
