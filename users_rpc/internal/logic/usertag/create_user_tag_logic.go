package usertag

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type CreateUserTagLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewCreateUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserTagLogic {
    return &CreateUserTagLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *CreateUserTagLogic) CreateUserTag(in *users.CreateUserTagReq) (*users.CreateUserTagResp, error) {

    now := time.Now().Unix()

    tag := &model.UserTag{
        Name:        in.Name,
        DisplayName: in.DisplayName,
        CreatedAt:   now,
        UpdatedAt:   now,
    }

    ret, err := l.svcCtx.UserTagModel.Insert(l.ctx, tag)
    if err != nil {
        return nil, err
    }

    id, err := ret.LastInsertId()
    if err != nil {
        return nil, err
    }

    return &users.CreateUserTagResp{
        Id: uint64(id),
    }, nil
}
