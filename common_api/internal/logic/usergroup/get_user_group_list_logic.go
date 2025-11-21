package usergroup

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupListLogic {
    return &GetUserGroupListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserGroupListLogic) GetUserGroupList(req *types.GetUserGroupListReq) (*types.GetUserGroupListResp, error) {

    rpcReq := &users.GetUserGroupListReq{
        Page:     req.Page,
        PageSize: req.PageSize,
    }

    rpcResp, err := l.svcCtx.UsersRpc.GetUserGroupList(l.ctx, rpcReq)
    if err != nil {
        return nil, err
    }

    resp := &types.GetUserGroupListResp{
        Total: rpcResp.Total,
        List:  make([]*types.UserGroupListItem, 0),
    }

    for _, g := range rpcResp.List {
        resp.List = append(resp.List, &types.UserGroupListItem{
            Id:          g.Id,
            Name:        g.Name,
            DisplayName: g.DisplayName,
            CreatedAt:   g.CreatedAt,
            UpdatedAt:   g.UpdatedAt,
        })
    }

    return resp, nil
}
