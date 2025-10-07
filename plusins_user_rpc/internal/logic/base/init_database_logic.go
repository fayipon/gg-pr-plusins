package base

import (
	"context"

	"github.com/fayipon/gg-pr-plusins/plusins_user_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/plusins_user_rpc/types/plusinsuserrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *plusinsuserrpc.Empty) (*plusinsuserrpc.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &plusinsuserrpc.BaseResp{}, nil
}
