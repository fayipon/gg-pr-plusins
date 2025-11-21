package userlevel

import (
    "net/http"

    "common_api/internal/logic/userlevel"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserLevelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.CreateUserLevelReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := userlevel.NewCreateUserLevelLogic(r.Context(), ctx)
        resp, err := l.CreateUserLevel(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
