package handler

import (
	"net/http"

	"common_api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/user",
		Handler: ctx.JwtMiddleware(GetUserHandler(ctx)),
	})

	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/user",
		Handler: ctx.JwtMiddleware(CreateUserHandler(ctx)),
	})

	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/user/list",
		Handler: ctx.JwtMiddleware(GetUserListHandler(ctx)),
	})

	server.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/login",
		Handler: LoginHandler(ctx),
	})

}
