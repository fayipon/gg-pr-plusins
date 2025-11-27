package userreferer

import (
    "net/http"

    "common_api/internal/logic/userreferer"
    "common_api/internal/response"
    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"
    "github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserRefererHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        var params types.GetUserRefererParam

        // 解析 GET URL Path /userreferer/:id
        if err := httpx.Parse(r, &params); err != nil {
            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, err.Error()),
            )
            return
        }

        if params.Id == 0 {
            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInvalidParams, "invalid id"),
            )
            return
        }

        req := types.GetUserRefererReq{
            Id: params.Id,
        }

        l := userreferer.NewGetUserRefererLogic(r.Context(), ctx)
        resp, err := l.GetUserReferer(&req)
        if err != nil {

            if e, ok := err.(*errorx.CodeError); ok {
                response.JsonError(w, r, e)
                return
            }

            response.JsonError(w, r,
                errorx.NewCodeError(r.Context(), errorx.ErrInternal, err.Error()),
            )
            return
        }

        response.OkJson(w, r, resp)
    }
}
