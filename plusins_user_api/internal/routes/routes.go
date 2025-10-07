package routes

import (
	"net/http"
	"plusins_user_api/internal/handler/user"
	"plusins_user_api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodPost,
				Path: "/user/login",
				Handler: user.LoginHandler(ctx),
			},
		},
	)
}