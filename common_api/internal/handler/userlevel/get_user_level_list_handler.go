package userlevel

import (
    "net/http"
    "strconv"

    "common_api/internal/logic/userlevel"
    "common_api/internal/response"
    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"
)

func GetUserLevelListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        // --- 手动读取 Query 参数 ---
        pageStr := r.URL.Query().Get("page")
        pageSizeStr := r.URL.Query().Get("page_size")

        // 默认值
        var page uint64 = 1
        var pageSize uint64 = 20

        if pageStr != "" {
            if v, err := strconv.ParseUint(pageStr, 10, 64); err == nil {
                page = v
            }
        }

        if pageSizeStr != "" {
            if v, err := strconv.ParseUint(pageSizeStr, 10, 64); err == nil {
                pageSize = v
            }
        }

        req := types.GetUserLevelListReq{
            Page:     int32(page),
            PageSize: int32(pageSize),
        }

        l := userlevel.NewGetUserLevelListLogic(r.Context(), ctx)
        resp, codeErr := l.GetUserLevelList(&req)
        if codeErr != nil {
            // ⭐ 强制类型断言为 *CodeError
            response.JsonError(w, r, codeErr.(*errorx.CodeError))
            return
        }

        response.OkJson(w, r, resp)
    }
}
