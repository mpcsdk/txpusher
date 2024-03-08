// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Txs is the golang structure of table txs for DAO operations like Where/Data.
type Txs struct {
	g.Meta    `orm:"table:txs, do:true"`
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
	UserId    interface{} //
}
