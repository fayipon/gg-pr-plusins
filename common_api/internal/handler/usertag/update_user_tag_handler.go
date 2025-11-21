package usertag

import (
    "net/http"

    "common_api/internal/logic/usertag"
    "common_api/internal/response"
    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"

    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserTagHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        // ⭐ 一次 Parse：同時解析 PathParam + JSON Body
        var req types.UpdateUserTagReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Errorf("Parse Error: %s", err.Error())
            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, err.Error()))
            return
        }

        // ⭐ PathParam id 必須 > 0
        if req.Id == 0 {
            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "invalid id"))
            return
        }

        // ⭐ Debug（可留可去）
        logx.Infof(">>> UpdateUserTag: id=%d name=%s display=%s",
            req.Id, req.Name, req.DisplayName)

        l := usertag.NewUpdateUserTagLogic(r.Context(), ctx)
        resp, codeErr := l.UpdateUserTag(&req)
        if codeErr != nil {
            response.JsonError(w, r, codeErr.(*errorx.CodeError))
            return
        }

        response.OkJson(w, r, resp)
    }
}
