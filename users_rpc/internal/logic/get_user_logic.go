package logic

import (
    "context"
    "database/sql"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
    return &GetUserLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserLogic) GetUser(in *users.GetUserReq) (*users.GetUserResp, error) {

    u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
    if err != nil {
        logx.Errorf("GetUser FindOne Error: id=%d, err=%+v", in.Id, err)
        return nil, err
    }

    resp := &users.GetUserResp{
        Id:        u.Id,
        Account:   u.Account,
        LevelId:   uint64(u.LevelId),
        GroupId:   uint64(u.GroupId),
        ParentId:  u.ParentId,
        RefererId: u.RefererId,
        Depth:     int64(u.Depth),
        Status:    int64(u.Status),
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
    }

    if u.EmailVerifiedAt.Valid {
        resp.EmailVerifiedAt = u.EmailVerifiedAt.Int64
    }
    if u.MobileVerifiedAt.Valid {
        resp.MobileVerifiedAt = u.MobileVerifiedAt.Int64
    }
    if u.KycVerifiedAt.Valid {
        resp.KycVerifiedAt = u.KycVerifiedAt.Int64
    }
    if u.ParentTree.Valid {
        resp.ParentTree = u.ParentTree.String
    }

    if u.LevelId > 0 {
        level, err := l.svcCtx.UserLevelModel.FindOne(l.ctx, u.LevelId)
        if err == nil {
            resp.UserLevel = &users.UserLevelInfo{
                Id:          level.Id,
                Name:        level.Name,
                DisplayName: level.DisplayName,
            }
        } else if err != sql.ErrNoRows {
            logx.Errorf("FindUserLevel error: %v", err)
        }
    }

    return resp, nil
}
