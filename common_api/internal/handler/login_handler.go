package handler

import (
    "net/http"

    "common_api/internal/utils/errorx"
    "common_api/internal/logic"
    "common_api/internal/response"
    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        // 解析请求
        var req types.LoginReq
        if err := httpx.Parse(r, &req); err != nil {
            response.JsonError(w, r, errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "en"))
            return
        }

        // 调用业务逻辑
        l := logic.NewLoginLogic(r.Context(), svcCtx)
        resp, codeErr := l.Login(&req)
        if codeErr != nil {
            response.JsonError(w, r, codeErr)
            return
        }

        // 成功响应
        response.OkJson(w, r, resp)

        /* 這邊抛错有包装过，后续以这为准 */
    }
}
