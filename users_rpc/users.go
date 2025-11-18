package main

import (
	"flag"
	"fmt"

	"users_rpc/internal/config"
	"users_rpc/internal/server"
	"users_rpc/internal/svc"

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
