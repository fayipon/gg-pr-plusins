package svc

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/plusins_merchant/internal/config"
    "github.com/fayipon/gg-pr-plusins/plusins_merchant/internal/ent"
)

type ServiceContext struct {
    Config config.Config
    DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
    client, err := ent.Open("postgres", c.CasbinDatabaseConf.GetDSN())
    if err != nil {
        panic(err)
    }

    if err := client.Schema.Create(context.Background()); err != nil {
        panic(err)
    }

    return &ServiceContext{
        Config: c,
        DB:     client,
    }
}
