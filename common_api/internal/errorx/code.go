package errorx

// 用户模块 10000 - 19999
const (
    ErrAccountNotFound = 10001
    ErrPasswordWrong   = 10002
    ErrUserNotFound    = 10003
)

// Token / Auth 20000 - 29999
const (
    ErrTokenMissing  = 20001
    ErrTokenInvalid  = 20002
    ErrTokenExpired  = 20003
)

// 参数错误 40000 - 49999
const (
    ErrInvalidParams = 40001
)

// 系统错误 50000 - 59999
const (
    ErrInternal = 50001
)
