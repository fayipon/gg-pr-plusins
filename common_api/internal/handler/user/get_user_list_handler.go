package handler

import (
    "net/http"

	"common_api/internal/svc"
    "common_api/internal/logic"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"

    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.UserListReq
        if err := httpx.Parse(r, &req); err != nil {
            errResp := errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "zh")
            httpx.WriteJson(w, http.StatusBadRequest, errResp)
            return
        }

        l := logic.NewGetUserListLogic(r.Context(), svcCtx)
        resp, err := l.GetUserList(&req)
        if err != nil {
            errResp := errorx.NewCodeError(r.Context(), errorx.ErrInternal, "zh")
            httpx.WriteJson(w, http.StatusBadRequest, errResp)
            return
        }

        httpx.OkJson(w, resp)
    }
}
