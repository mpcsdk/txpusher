// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TxsOffline is the golang structure of table txs_offline for DAO operations like Where/Data.
type TxsOffline struct {
	g.Meta      `orm:"table:txs_offline, do:true"`
	UserId      interface{} //
	ChainId     interface{} //
	BlockNumber interface{} //
}
