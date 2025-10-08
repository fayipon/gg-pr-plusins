package base

import (
	"context"

	"github.com/fayipon/gg-pr-plusins/plusins_api/internal/svc"
	"github.com/fayipon/gg-pr-plusins/plusins_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (*types.PingResp, error) {
	return &types.PingResp{
		Code: 0,
		Msg:  "plusins-api is alive",
	}, nil
}
