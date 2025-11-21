package main

import (
    "flag"
    "fmt"

    "common_api/internal/config"
    "common_api/internal/handler"
    "common_api/internal/svc"
	
	"github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/common_api.yaml", "config file")

func main() {
    flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)

    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()

    ctx := svc.NewServiceContext(c)
    handler.RegisterHandlers(server, ctx)

    fmt.Printf("Starting common-api at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
