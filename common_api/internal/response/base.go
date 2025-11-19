package response

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
)

type BaseResponse struct {
    Code          int         `json:"code"`
    Message       string      `json:"message"`
    LocaleMessage string      `json:"locale_message,omitempty"`
    Timestamp     int64       `json:"timestamp,omitempty"`
    Data          interface{} `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, r *http.Request, resp *BaseResponse) {
    httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, resp)
}
