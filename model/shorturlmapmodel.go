package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShortUrlMapModel = (*customShortUrlMapModel)(nil)

type (
	// ShortUrlMapModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortUrlMapModel.
	ShortUrlMapModel interface {
		shortUrlMapModel
	}

	customShortUrlMapModel struct {
		*defaultShortUrlMapModel
	}
)

// NewShortUrlMapModel returns a model for the database table.
func NewShortUrlMapModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShortUrlMapModel {
	return &customShortUrlMapModel{
		// 把 MySQL 连接和 Redis 配置一起交给 go-zero，后续查询会自动带上缓存能力。
		defaultShortUrlMapModel: newShortUrlMapModel(conn, c, opts...),
	}
}
