package handler

import (
    "net/http"

    user "common_api/internal/handler/user"
    userlevel "common_api/internal/handler/userlevel"
    usergroup "common_api/internal/handler/usergroup"
    usertag "common_api/internal/handler/usertag"
    "common_api/internal/svc"
    "github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {


    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/login",
        Handler: user.LoginHandler(ctx),
    })

	// --------------------
    // User 
    // --------------------
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
    
    // --------------------
    // User Level
    // --------------------
    // Create
    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/userlevel",
        Handler: ctx.JwtMiddleware(userlevel.CreateUserLevelHandler(ctx)),
    })

    // Get by ID: /userlevel/{id}
    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/userlevel/:id",
        Handler: ctx.JwtMiddleware(userlevel.GetUserLevelHandler(ctx)),
    })

    // List
    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/userlevel",
        Handler: ctx.JwtMiddleware(userlevel.GetUserLevelListHandler(ctx)),
    })

    // Update by ID: /userlevel/{id}
    server.AddRoute(rest.Route{
        Method:  http.MethodPut,
        Path:    "/userlevel/:id",
        Handler: ctx.JwtMiddleware(userlevel.UpdateUserLevelHandler(ctx)),
    })

    // Delete by ID: /userlevel/{id}
    server.AddRoute(rest.Route{
        Method:  http.MethodDelete,
        Path:    "/userlevel/:id",
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
        Path:    "/usergroup/:id",
        Handler: ctx.JwtMiddleware(usergroup.GetUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPut,
        Path:    "/usergroup/:id",
        Handler: ctx.JwtMiddleware(usergroup.UpdateUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodDelete,
        Path:    "/usergroup/:id",
        Handler: ctx.JwtMiddleware(usergroup.DeleteUserGroupHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/usergroup",
        Handler: ctx.JwtMiddleware(usergroup.GetUserGroupListHandler(ctx)),
    })

    
	// --------------------
    // User Tag
    // --------------------
    server.AddRoute(rest.Route{
        Method:  http.MethodPost,
        Path:    "/usertag",
        Handler: ctx.JwtMiddleware(usertag.CreateUserTagHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/usertag/:id",
        Handler: ctx.JwtMiddleware(usertag.GetUserTagHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodPut,
        Path:    "/usertag/:id",
        Handler: ctx.JwtMiddleware(usertag.UpdateUserTagHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodDelete,
        Path:    "/usertag/:id",
        Handler: ctx.JwtMiddleware(usertag.DeleteUserTagHandler(ctx)),
    })

    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/usertag",
        Handler: ctx.JwtMiddleware(usertag.GetUserTagListHandler(ctx)),
    })
}