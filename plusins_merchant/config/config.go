package config

import (
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

// Config 總設定結構
type Config struct {
    conf.RestConf
    CasbinDatabaseConf DatabaseConf
    CasbinConf          CasbinConf
}

// PostgreSQL 資料庫設定
type DatabaseConf struct {
    Type        string
    Host        string
    Port        int
    DBName      string
    Username    string
    Password    string
    SSLMode     string
    MaxOpenConn int
    CacheTime   int
}

// Casbin 模型設定
type CasbinConf struct {
    ModelText string
}

// DSN
func (c DatabaseConf) GetDSN() string {
    return sqlx.PostgresDataSourceName(sqlx.PostgresConf{
        Host:     c.Host,
        Port:     c.Port,
        User:     c.Username,
        Password: c.Password,
        DBName:   c.DBName,
        SSLMode:  c.SSLMode,
    })
}
