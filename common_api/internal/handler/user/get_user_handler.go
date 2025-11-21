package handler

import (
    "net/http"
    "strconv"

    "common_api/internal/logic"
    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"

    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        // 1. parse id
        idStr := r.URL.Query().Get("id")
        if idStr == "" {
            httpx.ErrorCtx(r.Context(), w, errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "en"))
            return
        }
        id, err := strconv.ParseUint(idStr, 10, 64)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "en"))
            return
        }

        // 2. call logic
        resp, err := logic.NewGetUserLogic(r.Context(), svcCtx).GetUser(&types.GetUserReq{Id: id})
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, errorx.NewCodeError(r.Context(), errorx.ErrInternal, "en"))
            return
        }

        // 3. success
        httpx.OkJsonCtx(r.Context(), w, resp)
    }
}
