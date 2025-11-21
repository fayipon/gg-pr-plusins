package logic

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
    return &GetUserListLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserListLogic) GetUserList(in *users.GetUserListReq) (*users.GetUserListResp, error) {

    page := in.Page
    pageSize := in.PageSize
    offset := (page - 1) * pageSize

    list, err := l.svcCtx.UserModel.List(l.ctx, int(offset), int(pageSize))
    if err != nil {
        return nil, err
    }

    total, err := l.svcCtx.UserModel.Count(l.ctx)
    if err != nil {
        return nil, err
    }

    resp := &users.GetUserListResp{
        Total: total,
    }

    for _, u := range list {
        resp.List = append(resp.List, &users.UserItem{
            Id:        u.Id,
            Account:   u.Account,
            Status:    u.Status,          
            LevelId:   u.LevelId,  
            ParentId:  u.ParentId,       
            GroupId:   u.GroupId,  
            CreatedAt: u.CreatedAt,       
        })
    }

    return resp, nil
}
