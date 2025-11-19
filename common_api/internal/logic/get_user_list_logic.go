package logic

import (
	"context"
	"common_api/internal/svc"
	"common_api/internal/types"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (*types.UserListResp, error) {

	rpcResp, err := l.svcCtx.UsersRpc.GetUserList(l.ctx, &users.GetUserListReq{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
		Keyword:  req.Keyword,
	})

	if err != nil {
		return nil, err
	}

	resp := &types.UserListResp{
		Total: rpcResp.Total,
		List:  make([]types.UserListItem, 0, len(rpcResp.List)),
	}

	for _, item := range rpcResp.List {
		resp.List = append(resp.List, types.UserListItem{
			Id:      item.Id,
			Account: item.Account,
		})
	}

	return resp, nil
}
