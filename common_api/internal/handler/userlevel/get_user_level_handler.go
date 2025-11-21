package userlevel

import (
    "errors"
    "net/http"
    "strconv"

    "common_api/internal/logic/userlevel"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserLevelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.GetUserLevelReq

        // 支援 GET Query ?id=xxx
        if r.Method == http.MethodGet {
            idStr := r.URL.Query().Get("id")
            if idStr != "" {
                if id, err := strconv.ParseUint(idStr, 10, 64); err == nil {
                    req.Id = id
                }
            }
        }

        // 支援 POST JSON body
        _ = httpx.Parse(r, &req)

        if req.Id == 0 {
            httpx.ErrorCtx(r.Context(), w, errors.New("id is required"))
            return
        }

        l := userlevel.NewGetUserLevelLogic(r.Context(), ctx)
        resp, err := l.GetUserLevel(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
