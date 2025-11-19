package logic

import (
    "context"
    "common_api/internal/errorx"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
    "google.golang.org/grpc/status"
)

type GetUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
    return &GetUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (*types.GetUserResp, error) {

    // 语言
    lang := "en"
    if v := l.ctx.Value("Lang"); v != nil {
        lang = v.(string)
    }

    // 调 RPC
    rpcResp, err := l.svcCtx.UsersRpc.GetUser(l.ctx, &users.GetUserReq{
        Id: req.Id,
    })

    // gRPC 错误处理
    if err != nil {
        st, ok := status.FromError(err)

        if ok && st.Code().String() == "NotFound" {
            return nil, errorx.NewCodeError(l.ctx, errorx.ErrUserNotFound, lang)
        }

        return nil, errorx.NewCodeError(l.ctx, errorx.ErrInternal, lang)
    }

    // 防止 nil 指针
    if rpcResp == nil {
        return nil, errorx.NewCodeError(l.ctx, errorx.ErrUserNotFound, lang)
    }

    // 返回
    return &types.GetUserResp{
        Id:        rpcResp.Id,
        Account:   rpcResp.Account,
        Status:    rpcResp.Status,
        LevelId:   rpcResp.LevelId,
        CreatedAt: rpcResp.CreatedAt,
    }, nil
}