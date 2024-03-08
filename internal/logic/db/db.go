package db

import (
	"txpusher/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

type sDB struct {
}

func new() *sDB {
	return &sDB{}
}
func init() {
	service.RegisterDB(new())
}
