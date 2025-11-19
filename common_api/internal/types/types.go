package types

// ----------------------
// Login
// ----------------------
type LoginReq struct {
    Account  string `json:"account"`
    Password string `json:"password"`
}

type LoginResp struct {
    Token string `json:"token"`
}

// ----------------------
// Create User
// ----------------------
type CreateUserReq struct {
    Account  string `json:"account"`
    Password string `json:"password"`
}

type CreateUserResp struct {
    Success bool `json:"success"`
}

// ----------------------
// Get User
// ----------------------
type GetUserReq struct {
    Id uint64 `json:"id"`
}

type GetUserResp struct {
    Id        uint64 `json:"id"`
    Account   string `json:"account"`
    Status    int64  `json:"status"`
    LevelId   int64  `json:"level_id"`
    CreatedAt int64  `json:"created_at"`
}

// ----------------------
// Filters
// ----------------------
type FilterItem struct {
    Field string      `json:"field"`
    Op    string      `json:"op"`
    Value interface{} `json:"value"`
}

// ----------------------
// User List
// ----------------------
type UserListReq struct {
    Page      int          `json:"page"`
    PageSize  int          `json:"page_size"`
    Keyword   string       `json:"keyword"`

    SortField string       `json:"sort_field"`
    SortOrder string       `json:"sort_order"`

    Filters []FilterItem `json:"filters"`
}

type UserListItem struct {
    Id        uint64 `json:"id"`
    Account   string `json:"account"`
    Status    int64  `json:"status"`
    LevelId   int64  `json:"level_id"`
    CreatedAt int64  `json:"created_at"`
}

type UserListResp struct {
    Total int64          `json:"total"`
    List  []UserListItem `json:"list"`
}
