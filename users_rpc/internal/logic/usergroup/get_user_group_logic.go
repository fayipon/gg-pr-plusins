package usergroup

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupLogic {
    return &GetUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserGroupLogic) GetUserGroup(in *users.GetUserGroupReq) (*users.GetUserGroupResp, error) {

    g, err := l.svcCtx.UserGroupModel.FindOne(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.GetUserGroupResp{
        Id:          g.Id,
        Name:        g.Name,
        DisplayName: g.DisplayName,
        Setting:     g.Setting,
        CreatedAt:   g.CreatedAt,
        UpdatedAt:   g.UpdatedAt,
    }, nil
}
