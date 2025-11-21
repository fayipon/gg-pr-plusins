package usergroup

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewUpdateUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserGroupLogic {
    return &UpdateUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *UpdateUserGroupLogic) UpdateUserGroup(in *users.UpdateUserGroupReq) (*users.UpdateUserGroupResp, error) {

    now := time.Now().Unix()

    data := &model.UserGroups{
        Id:          in.Id,
        Name:        in.Name,
        DisplayName: in.DisplayName,
        Setting:     in.Setting,
        UpdatedAt:   now,
    }

    err := l.svcCtx.UserGroupModel.Update(l.ctx, data)
    if err != nil {
        return nil, err
    }

    return &users.UpdateUserGroupResp{
        Success: true,
    }, nil
}
