package handler

import (
    "net/http"

    user "common_api/internal/handler/user"
    userlevel "common_api/internal/handler/userlevel"
    usergroup "common_api/internal/handler/usergroup"
    "common_api/internal/svc"
    "github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {

    //
    // User Routes
    //
    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/user",
        Handler: ctx.JwtMiddleware(user.CreateUserHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/user",
        Handler: ctx.JwtMiddleware(user.GetUserHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/user/list",
        Handler: ctx.JwtMiddleware(user.GetUserListHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/login",
        Handler: user.LoginHandler(ctx),
    })

    //
    // User Level Routes
    //
    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/userlevel",
        Handler: ctx.JwtMiddleware(userlevel.CreateUserLevelHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/userlevel",
        Handler: ctx.JwtMiddleware(userlevel.GetUserLevelHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/userlevel/list",
        Handler: ctx.JwtMiddleware(userlevel.GetUserLevelListHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/userlevel/update",
        Handler: ctx.JwtMiddleware(userlevel.UpdateUserLevelHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/userlevel/delete",
        Handler: ctx.JwtMiddleware(userlevel.DeleteUserLevelHandler(ctx)),
    })

	// --------------------
    // User Group
    // --------------------
    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/usergroup",
        Handler: ctx.JwtMiddleware(usergroup.CreateUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/usergroup",
        Handler: ctx.JwtMiddleware(usergroup.GetUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPut,
        Path:    "/usergroup",
        Handler: ctx.JwtMiddleware(usergroup.UpdateUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodDelete,
        Path:    "/usergroup",
        Handler: ctx.JwtMiddleware(usergroup.DeleteUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/usergroup/list",
        Handler: ctx.JwtMiddleware(usergroup.GetUserGroupListHandler(ctx)),
    })

}
