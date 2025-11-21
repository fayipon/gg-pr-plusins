package userlevel

import (
    "net/http"

    "common_api/internal/logic/userlevel"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserLevelListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.GetUserLevelListReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := userlevel.NewGetUserLevelListLogic(r.Context(), ctx)
        resp, err := l.GetUserLevelList(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
