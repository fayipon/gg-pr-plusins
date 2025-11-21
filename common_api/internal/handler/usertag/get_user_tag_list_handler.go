package usertag

import (
    "net/http"
    "strconv"

    "common_api/internal/logic/usertag"
    "common_api/internal/response"
    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"
)

func GetUserTagListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        pageStr := r.URL.Query().Get("page")
        pageSizeStr := r.URL.Query().Get("page_size")

        var page uint64 = 1
        var pageSize uint64 = 20

        if v, err := strconv.ParseUint(pageStr, 10, 64); err == nil {
            page = v
        }
        if v, err := strconv.ParseUint(pageSizeStr, 10, 64); err == nil {
            pageSize = v
        }

        req := types.GetUserTagListReq{
            Page:     int32(page),
            PageSize: int32(pageSize),
        }

        l := usertag.NewGetUserTagListLogic(r.Context(), ctx)
        resp, codeErr := l.GetUserTagList(&req)
        if codeErr != nil {
            response.JsonError(w, r, codeErr.(*errorx.CodeError))
            return
        }

        response.OkJson(w, r, resp)
    }
}
