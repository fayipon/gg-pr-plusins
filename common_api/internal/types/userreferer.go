package types

// ----------------------------
// UserReferer 基本结构
// ----------------------------
type CreateUserRefererReq struct {
    UserId      uint64 `json:"user_id"`
    ParentTree  string `json:"parent_tree"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
}

type CreateUserRefererResp struct {
    Id uint64 `json:"id"`
}

type GetUserRefererReq struct {
    Id uint64 `json:"id"`
}

type GetUserRefererParam struct {
    Id uint64 `path:"id"`
}

type UserRefererInfo struct {
    Id              uint64 `json:"id"`
    UserId          uint64 `json:"user_id"`
    ParentTree      string `json:"parent_tree"`
    Name            string `json:"name"`
    DisplayName     string `json:"display_name"`
    VisitCount      int64  `json:"visit_count"`
    RegisterCount   int64  `json:"register_count"`
    FirstDepositCount int64 `json:"first_deposit_count"`
    CreatedAt       int64  `json:"created_at"`
    UpdatedAt       int64  `json:"updated_at"`
}

type GetUserRefererResp struct {
    Info UserRefererInfo `json:"info"`
}

type UpdateUserRefererReq struct {
    Id          uint64 `json:"id" path:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    ParentTree  string `json:"parent_tree"`
}

type UpdateUserRefererResp struct {
    Success bool `json:"success"`
}

type DeleteUserRefererReq struct {
    Id uint64 `path:"id"`
}

type DeleteUserRefererResp struct {
    Success bool `json:"success"`
}

type GetUserRefererListReq struct {
    Page     int32 `json:"page,default=1"`
    PageSize int32 `json:"page_size,default=10"`
}

type GetUserRefererListResp struct {
    Total int64              `json:"total"`
    List  []UserRefererInfo  `json:"list"`
}
