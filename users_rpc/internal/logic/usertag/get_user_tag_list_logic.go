package usertag

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserTagListLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTagListLogic {
    return &GetUserTagListLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserTagListLogic) GetUserTagList(in *users.GetUserTagListReq) (*users.GetUserTagListResp, error) {

    page := in.Page
    pageSize := in.PageSize

    if page < 1 {
        page = 1
    }

    offset := (page - 1) * pageSize
    limit := pageSize

    list, err := l.svcCtx.UserTagModel.List(l.ctx, offset, limit)
    if err != nil {
        return nil, err
    }

    total, err := l.svcCtx.UserTagModel.Count(l.ctx)
    if err != nil {
        return nil, err
    }

    resp := &users.GetUserTagListResp{
        Total: total,
        List:  make([]*users.UserTagInfo, 0),
    }

    for _, t := range list {
        resp.List = append(resp.List, &users.UserTagInfo{
            Id:          t.Id,
            Name:        t.Name,
            DisplayName: t.DisplayName,
            CreatedAt:   t.CreatedAt,
            UpdatedAt:   t.UpdatedAt,
        })
    }

    return resp, nil
}
