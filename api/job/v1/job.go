package v1

import "github.com/gogf/gf/v2/frame/g"

type RecentTxsReq struct {
	g.Meta `path:"/RecentTxs" tags:"RecentTxs" method:"post" summary:"You first hello api"`
	Token  string `json:"token"`
}
type RecentTxsRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Seq    int         `json:"seq"`
	Data   interface{} `json:"data"`
}

// //
type TxsAckReq struct {
	g.Meta `path:"/TxsAc" tags:"TxsAc" method:"post" summary:"You first hello api"`
	Token  string `json:"token"`
	Seq    int    `json:"seq"`
}
type TxsAckRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
