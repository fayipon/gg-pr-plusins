package usergroup

import (
    "net/http"

    "common_api/internal/logic/usergroup"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUserGroupHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.DeleteUserGroupReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := usergroup.NewDeleteUserGroupLogic(r.Context(), ctx)
        resp, err := l.DeleteUserGroup(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
