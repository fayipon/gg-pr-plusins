package logic

import (
	"context"
	"common_api/internal/svc"
	"common_api/internal/types"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (*types.CreateUserResp, error) {

	rpcResp, err := l.svcCtx.UsersRpc.CreateUser(l.ctx, &users.CreateUserReq{
		Account:  req.Account,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.CreateUserResp{
		Id:      rpcResp.Id,
		Account: rpcResp.Account,
	}, nil
}
