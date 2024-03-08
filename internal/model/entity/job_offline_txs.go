// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// JobOfflineTxs is the golang structure for table job_offline_txs.
type JobOfflineTxs struct {
	ChainId   string `json:"chainId"   ` //
	Height    int    `json:"height"    ` //
	BlockHash string `json:"blockHash" ` //
	Ts        string `json:"ts"        ` //
	TxHash    string `json:"txHash"    ` //
	FromAddr  string `json:"fromAddr"  ` //
	ToAddr    string `json:"toAddr"    ` //
	Value     string `json:"value"     ` //
	Contract  string `json:"contract"  ` //
	Gas       string `json:"gas"       ` //
	GasPrice  string `json:"gasPrice"  ` //
	TxIdx     int    `json:"txIdx"     ` //
	LogIdx    int    `json:"logIdx"    ` //
	UserId    string `json:"userId"    ` //
}
