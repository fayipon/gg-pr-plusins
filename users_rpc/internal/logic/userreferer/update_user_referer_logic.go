package userreferer

import (
    "context"
    "time"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type UpdateUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUpdateUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRefererLogic {
    return &UpdateUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UpdateUserRefererLogic) UpdateUserReferer(req *users.UpdateUserRefererReq) (*users.UpdateUserRefererResp, error) {

    data := &model.UserReferer{
        Id:          req.Id,
        Name:        req.Name,
        DisplayName: req.DisplayName,
        ParentTree:  req.ParentTree,
        UpdatedAt:   time.Now().Unix(),
    }

    err := l.svcCtx.UserRefererModel.Update(l.ctx, data)
    if err != nil {
        return nil, err
    }

    return &users.UpdateUserRefererResp{
        Success: true,
    }, nil
}
