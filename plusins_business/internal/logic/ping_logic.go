package logic

import (
	"context"

	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() map[string]string {
	return map[string]string{
		"message": "pong from logic",
	}
}
