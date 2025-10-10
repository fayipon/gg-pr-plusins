package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/handler/base"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
)

// RegisterHandlers 注册所有 handler
func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  "GET",
				Path:    "/ping",
				Handler: base.PingHandler(ctx), // ✅ 调用 base/ping_handler.go
			},
		},
	)
}
