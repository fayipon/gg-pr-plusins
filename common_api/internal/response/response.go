package response

import (
    "net/http"

    "common_api/internal/errorx"
    "github.com/zeromicro/go-zero/rest/httpx"
)

type SuccessResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// 成功响应
func OkJson(w http.ResponseWriter, r *http.Request, data interface{}) {
    httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, &SuccessResponse{
        Code:    0,
        Message: "Success",
        Data:    data,
    })
}

// 错误响应
func JsonError(w http.ResponseWriter, r *http.Request, err *errorx.CodeError) {
    httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, err)
}
