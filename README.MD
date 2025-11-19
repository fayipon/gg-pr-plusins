
目錄結構 
│  docker-compose.yml
│  README.MD
│
├─common_api
│  │  api.go
│  │  go.mod
│  │  go.sum
│  │
│  ├─etc
│  │      common_api.yaml
│  │
│  └─internal
│      ├─config
│      │      config.go
│      │
│      ├─handler
│      │      create_user_handler.go
│      │      get_user_handler.go
│      │      get_user_list_handler.go
│      │      routers.go
│      │
│      ├─logic
│      │      create_user_logic.go
│      │      get_user_list_logic.go
│      │      get_user_logic.go
│      │
│      ├─svc
│      │      service_context.go
│      │
│      └─types
│              types.go
│
└─users_rpc
    │  .gitignore
    │  Dockerfile
    │  go.mod
    │  go.sum
    │  Makefile
    │  users.go
    │  users.proto
    │
    ├─etc
    │      users.yaml
    │
    ├─internal
    │  ├─config
    │  │      config.go
    │  │
    │  ├─logic
    │  │      create_user_logic.go
    │  │      get_user_list_logic.go
    │  │      get_user_logic.go
    │  │
    │  ├─model
    │  │      users_model.go
    │  │
    │  ├─server
    │  │      users_server.go
    │  │
    │  └─svc
    │          service_context.go
    │
    └─users
            users.pb.go
            users_grpc.pb.go


common_api , 主負責 API層處理
*_rpc , rpc 層 , 以數據庫表為單位的 數據處理


------------------
Docker 

TODO 

------------------
PORT LIST
COMMON_API　8080

USERS_RPC  8081