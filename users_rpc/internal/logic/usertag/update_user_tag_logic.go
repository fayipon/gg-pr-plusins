package usertag

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserTagLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewUpdateUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserTagLogic {
    return &UpdateUserTagLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *UpdateUserTagLogic) UpdateUserTag(in *users.UpdateUserTagReq) (*users.UpdateUserTagResp, error) {

    tag := &model.UserTag{
        Id:          in.Id,
        Name:        in.Name,
        DisplayName: in.DisplayName,
        UpdatedAt:   time.Now().Unix(),
    }

    err := l.svcCtx.UserTagModel.Update(l.ctx, tag)
    if err != nil {
        return nil, err
    }

    return &users.UpdateUserTagResp{
        Success: true,
    }, nil
}
