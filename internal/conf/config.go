package conf

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type Cache struct {
	ApiInterval              int `json:"apiInterval" v:"required|min:1"`
	SessionDuration          int `json:"sessionDuration" v:"required|min:1"`
	DBCacheDuration          int `json:"dbCacheDuration" v:"required|min:1"`
	VerificationCodeDuration int `json:"VerificationCodeDuration" v:"required|min:1"`
	LimitSendPhoneCount      int `json:"limitSendPhoneCount" v:"required|min:1"`
	LimitSendPhoneDuration   int `json:"limitSendPhoneDuration" v:"required|min:1"`
	LimitSendMailCount       int `json:"limitSendMailCount" v:"required|min:1"`
	LimitSendMailDuration    int `json:"limitSendMailDuration" v:"required|min:1"`
}

type Server struct {
	Address string `json:"address" v:"required"`
	WorkId  int    `json:"workId" v:"required|min:1"`
	Name    string `json:"name" v:"required"`
}

type Nrpcfg struct {
	NatsUrl string `json:"natsUrl" v:"required"`
}

// //
type Cfg struct {
	Server    *Server `json:"server" v:"required"`
	Cache     *Cache  `json:"cache" v:"required"`
	JaegerUrl string  `json:"jaegerUrl" v:"required"`
	Nrpc      *Nrpcfg `json:"nrpc" v:"required"`
	Txdatarpc string  `json:"txdatarpc" v:"required"`
}

var Config = &Cfg{}

func init() {
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()
	v, err := cfg.Data(ctx)
	if err != nil {
		panic(err)
	}
	val := gvar.New(v)
	err = val.Structs(Config)
	if err != nil {
		panic(err)
	}
	if err := g.Validator().Data(Config).Run(ctx); err != nil {
		panic(err)
	}
}
