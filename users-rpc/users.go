package main

import (
	"flag"
	"fmt"

	"users-rpc/internal/config"
	"users-rpc/internal/server"
	"users-rpc/internal/svc"
	"users-rpc/users"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := service.NewService(
		service.WithName(c.Name),
		service.WithListenOn(c.ListenOn),
	)

	users.RegisterUsersServer(s.Server(), server.NewUsersServer(ctx))

	fmt.Printf("Starting RPC server at %s...\n", c.ListenOn)
	s.Start()
}
