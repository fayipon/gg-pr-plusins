package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/handler"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
)

var configFile = flag.String("f", "etc/plusins_business.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting plusins-business-api at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
