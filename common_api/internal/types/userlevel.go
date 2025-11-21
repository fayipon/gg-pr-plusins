package types

type CreateUserLevelReq struct {
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    Setting     string `json:"setting"`
}

type CreateUserLevelResp struct {
    Id uint64 `json:"id"`
}

type UpdateUserLevelReq struct {
    Id          uint64 `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    Setting     string `json:"setting"`
}

type UpdateUserLevelResp struct {
    Success bool `json:"success"`
}

type DeleteUserLevelParam struct {
    Id uint64 `path:"id"`
}

type DeleteUserLevelReq struct {
    Id uint64 `json:"id"`
}

type DeleteUserLevelResp struct {
    Success bool `json:"success"`
}

type GetUserLevelParam struct {
    Id uint64 `path:"id"`
}

type GetUserLevelReq struct {
    Id uint64 `json:"id"`
}

type UserLevelInfo struct {
    Id          uint64 `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    Setting     string `json:"setting"`
    CreatedAt   int64  `json:"created_at"`
    UpdatedAt   int64  `json:"updated_at"`
}

type GetUserLevelResp = UserLevelInfo

type GetUserLevelListReq struct {
    Page     int32 `json:"page"`
    PageSize int32 `json:"page_size"`
}

type UserLevelListItem = UserLevelInfo

type GetUserLevelListResp struct {
    Total int64                `json:"total"`
    List  []*UserLevelListItem `json:"list"`
}
