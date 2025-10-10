gg-pr-plusins/
├── docker-compose.yml
├── plusins_api/
├── plusins_business/
├── plusins_rpc/
└── plusins_merchant/
    ├── main.go
    ├── etc/
    │   └── plusins_merchant.yaml
    ├── internal/
    │   ├── config/
    │   │   └── config.go
    │   ├── ent/
    │   │   └── schema/
    │   │       └── merchant.go
    │   ├── handler/
    │   │   ├── merchant/
    │   │   │   ├── create_merchant_handler.go
    │   │   │   ├── list_merchant_handler.go
    │   │   │   ├── get_merchant_handler.go
    │   │   │   ├── update_merchant_handler.go
    │   │   │   └── delete_merchant_handler.go
    │   │   └── routes.go
    │   ├── logic/
    │   │   └── merchant/
    │   │       ├── create_merchant_logic.go
    │   │       ├── list_merchant_logic.go
    │   │       ├── get_merchant_logic.go
    │   │       ├── update_merchant_logic.go
    │   │       └── delete_merchant_logic.go
    │   ├── svc/
    │   │   └── service_context.go
    │   └── types/
    │       └── merchant.go
    └── Dockerfile
