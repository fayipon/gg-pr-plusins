package main

import (
	"flag"
	"fmt"

	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/config"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/server"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		users.RegisterUsersServer(grpcServer, server.NewUsersServer(ctx))
	})

	defer s.Stop()

	fmt.Printf("Starting users-rpc at %s...\n", c.ListenOn)
	s.Start()
}
