package db

import (
	"time"
	"txpusher/internal/conf"
	"txpusher/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

type sDB struct {
	dbDuration time.Duration
}

func new() *sDB {
	return &sDB{
		dbDuration: time.Duration(conf.Config.Cache.DBCacheDuration) * time.Second,
	}
}
func init() {
	service.RegisterDB(new())
}
