// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// JobOfflineTxs is the golang structure of table job_offline_txs for DAO operations like Where/Data.
type JobOfflineTxs struct {
	g.Meta    `orm:"table:job_offline_txs, do:true"`
	ChainId   interface{} //
	Height    interface{} //
	BlockHash interface{} //
	Ts        interface{} //
	TxHash    interface{} //
	FromAddr  interface{} //
	ToAddr    interface{} //
	Value     interface{} //
	Contract  interface{} //
	Gas       interface{} //
	GasPrice  interface{} //
	TxIdx     interface{} //
	LogIdx    interface{} //
	UserId    interface{} //
}
