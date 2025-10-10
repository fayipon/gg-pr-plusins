package handler

import (
	"net/http"

	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// RegisterHandlers 注册所有路由（示例模板）
func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	// 示例：添加一个健康检查 API
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ping",
		Handler: pingHandler(ctx),
	})
}

// pingHandler 用于测试服务是否启动成功
func pingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, map[string]interface{}{
			"message": "plusins-business running",
		})
	}
}
