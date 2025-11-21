package usertag

import (
    "net/http"

    "common_api/internal/logic/usertag"
    "common_api/internal/response"
    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"

    "github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteUserTagHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        // ⭐ RESTful Path Param (与 UserGroup/UserLevel 同风格)
        var params types.DeleteUserTagParam
        if err := httpx.Parse(r, &params); err != nil {
            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, err.Error()))
            return
        }

        if params.Id == 0 {
            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "invalid id"))
            return
        }

        req := types.DeleteUserTagReq{
            Id: params.Id,
        }

        l := usertag.NewDeleteUserTagLogic(r.Context(), ctx)
        resp, codeErr := l.DeleteUserTag(&req)
        if codeErr != nil {
            response.JsonError(w, r, codeErr.(*errorx.CodeError))
            return
        }

        response.OkJson(w, r, resp)
    }
}
