// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"errors"

	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/model"

	"github.com/zeromicro/go-zero/core/logx"
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
	// 查看短链接，输入 q1mi.cn/lusytc -> 重定向到真实的链接
	// req.ShortUrl = lusytc
	// 1. 根据短链接查询原始的长链接
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
