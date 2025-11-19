package handler

import (
    "net/http"
    "common_api/internal/logic"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.GetUserReq

        // 对 GET 请求，必须使用 ParseForm
        if err := httpx.ParseForm(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := logic.NewGetUserLogic(r.Context(), svcCtx)
        resp, err := l.GetUser(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
