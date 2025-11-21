package usergroup

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
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

func (l *GetUserGroupListLogic) GetUserGroupList(in *users.GetUserGroupListReq) (*users.GetUserGroupListResp, error) {

    page := in.Page
    pageSize := in.PageSize
    offset := (page - 1) * pageSize

    list, err := l.svcCtx.UserGroupModel.List(l.ctx, offset, pageSize)
    if err != nil {
        return nil, err
    }

    total, err := l.svcCtx.UserGroupModel.Count(l.ctx)
    if err != nil {
        return nil, err
    }

    resp := &users.GetUserGroupListResp{
        Total: total,
        List:  make([]*users.UserGroupInfo, 0),
    }

    for _, g := range list {
        resp.List = append(resp.List, &users.UserGroupInfo{
            Id:          g.Id,
            Name:        g.Name,
            DisplayName: g.DisplayName,
        })
    }

    return resp, nil
}
