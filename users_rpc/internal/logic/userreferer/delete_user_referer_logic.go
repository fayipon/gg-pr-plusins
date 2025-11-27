package userreferer

import (
    "context"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type DeleteUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewDeleteUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRefererLogic {
    return &DeleteUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *DeleteUserRefererLogic) DeleteUserReferer(req *users.DeleteUserRefererReq) (*users.DeleteUserRefererResp, error) {

    err := l.svcCtx.UserRefererModel.Delete(l.ctx, req.Id)
    if err != nil {
        return nil, err
    }

    return &users.DeleteUserRefererResp{
        Success: true,
    }, nil
}
