package types

//
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

//
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

//
// ----------------------
// Get User
// ----------------------
type GetUserReq struct {
    Id uint64 `json:"id" form:"id"`
}

type GetUserResp struct {
    Id              uint64         `json:"id"`
    Account         string         `json:"account"`
    LevelId         uint64          `json:"level_id"`
    GroupId         uint64         `json:"group_id"`
    EmailVerifiedAt int64          `json:"email_verified_at"`
    MobileVerifiedAt int64         `json:"mobile_verified_at"`
    KycVerifiedAt   int64          `json:"kyc_verified_at"`
    ParentId        uint64         `json:"parent_id"`
    ParentTree      string         `json:"parent_tree"`
    Depth           int64          `json:"depth"`
    RefererId       uint64         `json:"referer_id"`
    Status          int64          `json:"status"`
    CreatedAt       int64          `json:"created_at"`
    UpdatedAt       int64          `json:"updated_at"`
    UserLevel       *UserLevelInfo `json:"user_level,omitempty"`
}

//
// User Level
//
type UserLevelInfo struct {
    Id          uint64  `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
}
//
// ----------------------
// User List
// ----------------------
type UserListReq struct {
    Page     int    `json:"page"`
    PageSize int    `json:"page_size"`
    Keyword  string `json:"keyword"`
}

type UserListItem struct {
    Id        uint64 `json:"id"`
    Account   string `json:"account"`
    Status    int64  `json:"status"`
    LevelId   uint64 `json:"level_id"`
    CreatedAt int64  `json:"created_at"`

    UserLevel *UserLevelInfo `json:"user_level,omitempty"`
}

type UserListResp struct {
    Total int64          `json:"total"`
    List  []UserListItem `json:"list"`
}
