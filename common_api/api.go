package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	
	"common_api/internal/config"
	"common_api/internal/handler"
	"common_api/internal/svc"
)

var configFile = flag.String("f", "etc/common_api.yaml", "api config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting common_api at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
