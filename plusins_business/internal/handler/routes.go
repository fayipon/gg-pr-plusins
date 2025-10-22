package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/handler/base"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
)

// RegisterHandlers 注册所有 handler
func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	server.AddRoutes([]rest.Route{
			{Method:  "GET", Path:    "/ping", Handler: base.PingHandler(ctx)},
			{Method: http.MethodPost, Path: "/merchant", Handler: merchant.CreateMerchantHandler(ctx)},
			{Method: http.MethodGet, Path: "/merchant/list", Handler: merchant.ListMerchantHandler(ctx)},
			{Method: http.MethodGet, Path: "/merchant/:id", Handler: merchant.GetMerchantHandler(ctx)},
			{Method: http.MethodPut, Path: "/merchant/:id", Handler: merchant.UpdateMerchantHandler(ctx)},
			{Method: http.MethodDelete, Path: "/merchant/:id", Handler: merchant.DeleteMerchantHandler(ctx)},
		},
	)
}
