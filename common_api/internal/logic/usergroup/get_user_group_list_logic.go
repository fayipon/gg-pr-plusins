package usergroup

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupListLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupListLogic {
    return &GetUserGroupListLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserGroupListLogic) GetUserGroupList(req *types.GetUserGroupListReq) (*types.GetUserGroupListResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUserGroupList(l.ctx, &users.GetUserGroupListReq{
        Page:     req.Page,
        PageSize: req.PageSize,
    })

    if err != nil {
        return nil, err
    }

    resp := &types.GetUserGroupListResp{
        Total: rpcResp.Total,
        List:  make([]types.UserGroupListItem, 0),
    }

    for _, g := range rpcResp.List {
        resp.List = append(resp.List, types.UserGroupListItem{
            Id:          g.Id,
            Name:        g.Name,
            DisplayName: g.DisplayName,
        })
    }

    return resp, nil
}
