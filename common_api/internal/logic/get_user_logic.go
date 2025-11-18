package logic

import (
	"context"

	"common_api/internal/svc"
	"common_api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (*types.GetUserResp, error) {

	rpcResp, err := l.svcCtx.UsersRpc.GetUser(l.ctx, &users.GetUserReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetUserResp{
		Id:      rpcResp.Id,
		Account: rpcResp.Account,
	}, nil
}
