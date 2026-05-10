// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/model"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 查看短链接，输入 qlmi.cn/lusytc -> 重定向到真实的链接
	// req.ShortUrl = lusytc
	// 1. 根据短链接查询原始的长链接
	// 1.0 布隆过滤器
	// 不存在的短链接直接返回404，不需要后续处理
	// a. 基于内存版本，服务重启之后就没了，所以每次重启都要重新加载一下已有的短链接

	// b. 基于Redis版本,go-zero的布隆过滤器是基于Redis实现的，重启之后数据还在，不需要重新加载。
	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("Bloom filter Exists failed", logx.LogField{Key: "shortUrl", Value: req.ShortUrl}, logx.LogField{Key: "err", Value: err.Error()})
	}
	if !exist {
		return nil, Err404
	}
	fmt.Println("开始查询缓存和DB...")
	// 1.1 查询数据库之前可增加缓存层
	// go-zero的缓存支持singleflight
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(
		l.ctx,
		sql.NullString{String: req.ShortUrl, Valid: true},
	)
	if err != nil {
		if err == model.ErrNotFound {
			// 未找到对应记录时返回 404，交给 handler 层输出错误响应。
			return nil, errors.New("404")
		}

		logx.Errorw("ShortUrlModel.FindOneBySurl failed",
			logx.LogField{Key: "shortUrl", Value: req.ShortUrl},
			logx.LogField{Key: "err", Value: err.Error()},
		)
		return nil, err
	}

	// 把查到的长链接返回给 handler，由 handler 执行重定向。
	return &types.ShowResponse{
		LongUrl: u.Lurl.String,
	}, nil
}
