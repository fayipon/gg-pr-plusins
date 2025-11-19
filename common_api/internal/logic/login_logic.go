package logic

import (
    "context"
    "time"

    "common_api/internal/errorx"
    "common_api/internal/svc"
    "common_api/internal/types"

    "github.com/golang-jwt/jwt/v4"
    "github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
    return &LoginLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, *errorx.CodeError) {

    // 获取语言 (LangMiddleware 已放入 ctx)
    lang := "en"
    if v := l.ctx.Value("Lang"); v != nil {
        if s, ok := v.(string); ok {
            lang = s
        }
    }

    // 校验账号
    if req.Account != "admin" {
        return nil, errorx.NewCodeError(
            l.ctx,
            errorx.ErrAccountNotFound,
            lang,
        )
    }

    // 校验密码
    if req.Password != "123456" {
        return nil, errorx.NewCodeError(
            l.ctx,
            errorx.ErrPasswordWrong,
            lang,
        )
    }

    // 生成 token
    now := time.Now().Unix()
    expire := now + l.svcCtx.Config.JwtAuth.AccessExpire

    claims := jwt.MapClaims{
        "account": req.Account,
        "iat":     now,
        "exp":     expire,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenStr, err := token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
    if err != nil {
        return nil, errorx.NewCodeError(
            l.ctx,
            errorx.ErrInternal,
            lang,
        )
    }

    return &types.LoginResp{
        Token: tokenStr,
    }, nil
}
