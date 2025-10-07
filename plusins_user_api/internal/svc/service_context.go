package svc

import (
	"plusins_user_api/internal/config"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/token"
	"plusins_user_api/internal/middleware"
)


type ServiceContext struct {
	Config config.Config
	JwtAuth *middleware.JwtAuth
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config: c,
        JwtAuth: middleware.NewJwtAuth(c.Auth.AccessSecret),
    }
}
