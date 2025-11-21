package userlevel

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLevelListLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLevelListLogic {
    return &GetUserLevelListLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserLevelListLogic) GetUserLevelList(in *users.GetUserLevelListReq) (*users.GetUserLevelListResp, error) {

    page := in.Page
    pageSize := in.PageSize
    offset := (page - 1) * pageSize

    // 列表
    list, err := l.svcCtx.UserLevelModel.List(l.ctx, offset, pageSize)
    if err != nil {
        return nil, err
    }

    // 总数
    total, err := l.svcCtx.UserLevelModel.Count(l.ctx)
    if err != nil {
        return nil, err
    }

    resp := &users.GetUserLevelListResp{
        Total: total,
        List:  make([]*users.UserLevelInfo, 0),
    }

    for _, lv := range list {
        resp.List = append(resp.List, &users.UserLevelInfo{
            Id:          lv.Id,
            Name:        lv.Name,
            DisplayName: lv.DisplayName,
        })
    }

    return resp, nil
}
