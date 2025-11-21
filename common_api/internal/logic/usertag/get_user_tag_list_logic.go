package usertag

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserTagListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTagListLogic {
    return &GetUserTagListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserTagListLogic) GetUserTagList(req *types.GetUserTagListReq) (*types.GetUserTagListResp, error) {

    rpcReq := &users.GetUserTagListReq{
        Page:     req.Page,
        PageSize: req.PageSize,
    }

    rpcResp, err := l.svcCtx.UsersRpc.GetUserTagList(l.ctx, rpcReq)
    if err != nil {
        return nil, err
    }

    resp := &types.GetUserTagListResp{
        Total: rpcResp.Total,
        List:  make([]*types.UserTagListItem, 0),
    }

    for _, v := range rpcResp.List {
        resp.List = append(resp.List, &types.UserTagListItem{
            Id:          v.Id,
            Name:        v.Name,
            DisplayName: v.DisplayName,
            CreatedAt:   v.CreatedAt,
            UpdatedAt:   v.UpdatedAt,
        })
    }

    return resp, nil
}
