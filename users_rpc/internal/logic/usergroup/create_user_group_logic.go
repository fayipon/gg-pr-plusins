package usergroup

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type CreateUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewCreateUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserGroupLogic {
    return &CreateUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *CreateUserGroupLogic) CreateUserGroup(in *users.CreateUserGroupReq) (*users.CreateUserGroupResp, error) {

    now := time.Now().Unix()

    data := &model.UserGroups{
        Name:        in.Name,
        DisplayName: in.DisplayName,
        Setting:     in.Setting,
        CreatedAt:   now,
        UpdatedAt:   now,
    }

    res, err := l.svcCtx.UserGroupModel.Insert(l.ctx, data)
    if err != nil {
        return nil, err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return nil, err
    }

    return &users.CreateUserGroupResp{
        Id: uint64(id),
    }, nil
}
