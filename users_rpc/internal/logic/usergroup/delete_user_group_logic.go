package usergroup

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewDeleteUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserGroupLogic {
    return &DeleteUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *DeleteUserGroupLogic) DeleteUserGroup(in *users.DeleteUserGroupReq) (*users.DeleteUserGroupResp, error) {

    err := l.svcCtx.UserGroupModel.Delete(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.DeleteUserGroupResp{
        Success: true,
    }, nil
}
