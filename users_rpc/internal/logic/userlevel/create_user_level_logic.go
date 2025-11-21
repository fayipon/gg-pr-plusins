package userlevel

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLevelLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewCreateUserLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLevelLogic {
    return &CreateUserLevelLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *CreateUserLevelLogic) CreateUserLevel(in *users.CreateUserLevelReq) (*users.CreateUserLevelResp, error) {

    now := time.Now().Unix()

    data := &model.UserLevels{
        Name:        in.Name,
        DisplayName: in.DisplayName,
        Setting:     in.Setting,
        CreatedAt:   now,
        UpdatedAt:   now,
    }

    res, err := l.svcCtx.UserLevelModel.Insert(l.ctx, data)
    if err != nil {
        return nil, err
    }

    id, _ := res.LastInsertId()

    return &users.CreateUserLevelResp{
        Id: uint64(id),
    }, nil
}