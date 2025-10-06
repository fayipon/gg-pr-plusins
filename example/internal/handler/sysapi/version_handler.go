package sysapi

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
    "github.com/fayipon/gg-pr-plusins/example/internal/logic/sysapi"
    "github.com/fayipon/gg-pr-plusins/example/internal/svc"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        l := sysapi.NewVersionLogic(r.Context(), svcCtx)
        resp, err := l.Version()
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }
        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
