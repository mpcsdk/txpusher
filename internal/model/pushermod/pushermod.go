package pushermod

import "txpusher/internal/model/entity"

type Tansaction struct {
	ChainId   string `json:"chainid"`
	Height    int    `json:"height"`
	Blockhash string `json:"blockhash"`
	Ts        string `json:"timestamp"`
	Txhash    string `json:"txhash"`
	Toaddr    string `json:"toaddr"`
	Fromaddr  string `json:"fromaddr"`
	Value     string `json:"value"`
	Contract  string `json:"contract"`
	Gas       string `json:"gas"`
	Gasprice  string `json:"gasprice"`
	TxIdx     int    `json:"txidx"`
	LogIdx    int    `json:"logidx"`
}

func Trans2Tx(transaction *Tansaction, userId string) *entity.Txs {
	return &entity.Txs{
		ChainId:   transaction.ChainId,
		Height:    transaction.Height,
		BlockHash: transaction.Blockhash,
		Ts:        transaction.Ts,
		TxHash:    transaction.Txhash,
		ToAddr:    transaction.Toaddr,
		FromAddr:  transaction.Fromaddr,
		Value:     transaction.Value,
		Contract:  transaction.Contract,

		Gas:      transaction.Gas,
		GasPrice: transaction.Gasprice,
		UserId:   userId,
	}
}
func Trans2Txs(trans []*Tansaction) []*entity.Txs {
	data := []*entity.Txs{}
	for _, t := range trans {
		data = append(data, Trans2Tx(t, ""))
	}

	return data
}
