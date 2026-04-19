// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	ShortUrlModel     model.ShortUrlMapModel
	Sequence          sequence.Sequence
	ShortUrlBlackList map[string]struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)

	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}

	return &ServiceContext{
		Config: c,
		// goctl 不会自动改 svc，这里必须手动把 CacheRedis 传给 model，
		// model 里的 CachedConn 才会真正启用 Redis 缓存。
		ShortUrlModel:     model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:          sequence.NewMySQL(c.Sequence.DSN),
		ShortUrlBlackList: m,
	}
}
