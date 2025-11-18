package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"

	"users_rpc/internal/config"
	"users_rpc/internal/server"
	"users_rpc/internal/svc"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *service.ServiceGroup) {
		server.RegisterUsersServer(grpcServer, ctx)
	})
	defer s.Stop()

	fmt.Printf("Starting Users RPC Server at %s...\n", c.ListenOn)
	s.Start()
}
