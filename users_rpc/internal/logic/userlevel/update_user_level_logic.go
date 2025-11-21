package userlevel

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLevelLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewUpdateUserLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLevelLogic {
    return &UpdateUserLevelLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *UpdateUserLevelLogic) UpdateUserLevel(in *users.UpdateUserLevelReq) (*users.UpdateUserLevelResp, error) {

    now := time.Now().Unix()

    data := &model.UserLevels{
        Id:          in.Id,
        Name:        in.Name,
        DisplayName: in.DisplayName,
        Setting:     in.Setting,
        UpdatedAt:   now,
    }

    err := l.svcCtx.UserLevelModel.Update(l.ctx, data)
    if err != nil {
        return nil, err
    }

    return &users.UpdateUserLevelResp{
        Success: true,
    }, nil
}
