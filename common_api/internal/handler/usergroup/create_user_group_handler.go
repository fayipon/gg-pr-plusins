package usergroup

import (
    "net/http"

    "common_api/internal/logic/usergroup"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserGroupHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.CreateUserGroupReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := usergroup.NewCreateUserGroupLogic(r.Context(), ctx)
        resp, err := l.CreateUserGroup(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
