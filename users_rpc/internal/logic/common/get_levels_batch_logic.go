package common

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetLevelsBatchLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetLevelsBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLevelsBatchLogic {
    return &GetLevelsBatchLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetLevelsBatchLogic) GetLevelsBatch(in *users.LevelBatchReq) (*users.LevelBatchResp, error) {

    if len(in.Ids) == 0 {
        return &users.LevelBatchResp{List: []*users.UserLevelInfo{}}, nil
    }

    levels, err := l.svcCtx.UserLevelModel.FindByIds(l.ctx, in.Ids)
    if err != nil {
        logx.Errorf("FindByIds error: %v", err)
        return nil, err
    }

    resp := make([]*users.UserLevelInfo, 0, len(levels))
    for _, lv := range levels {
        resp = append(resp, &users.UserLevelInfo{
            Id:          lv.Id,
            Name:        lv.Name,
            DisplayName: lv.DisplayName,
        })
    }

    return &users.LevelBatchResp{
        List: resp,
    }, nil
}
