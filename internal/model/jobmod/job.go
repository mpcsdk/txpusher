package jobmod

type Tansaction struct {
	ChainId   uint64 `json:"chainid"`
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
