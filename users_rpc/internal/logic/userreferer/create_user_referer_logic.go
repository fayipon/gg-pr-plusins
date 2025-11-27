package userreferer

import (
    "context"
    "time"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"

    // ⭐ proto types
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type CreateUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewCreateUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserRefererLogic {
    return &CreateUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *CreateUserRefererLogic) CreateUserReferer(req *users.CreateUserRefererReq) (*users.CreateUserRefererResp, error) {

    data := &model.UserReferer{
        UserId:            req.UserId,
        ParentTree:        req.ParentTree,
        Name:              req.Name,
        DisplayName:       req.DisplayName,
        VisitCount:        0,
        RegisterCount:     0,
        FirstDepositCount: 0,
        CreatedAt:         time.Now().Unix(),
        UpdatedAt:         time.Now().Unix(),
    }

    res, err := l.svcCtx.UserRefererModel.Insert(l.ctx, data)
    if err != nil {
        return nil, err
    }

    id, _ := res.LastInsertId()
    return &users.CreateUserRefererResp{
        Id: uint64(id),
    }, nil
}
