package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
)

// PingHandler 测试接口
func PingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, map[string]string{
			"message": "pong",
		})
	}
}
