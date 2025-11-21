package types


// -----------------------
// RESTful Path Params
// -----------------------
type GetUserGroupParam struct {
    Id uint64 `path:"id"`
}

type UpdateUserGroupParam struct {
    Id uint64 `path:"id"`
}

type DeleteUserGroupParam struct {
    Id uint64 `path:"id"`
}

type CreateUserGroupReq struct {
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    Setting     string `json:"setting"`
}

type CreateUserGroupResp struct {
    Id uint64 `json:"id"`
}

type UpdateUserGroupReq struct {
    Id          uint64 `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    Setting     string `json:"setting"`
}

type UpdateUserGroupResp struct {
    Success bool `json:"success"`
}

type DeleteUserGroupReq struct {
    Id uint64 `json:"id" form:"id"`
}

type DeleteUserGroupResp struct {
    Success bool `json:"success"`
}

type GetUserGroupReq struct {
    Id uint64 `json:"id" form:"id"`
}

type UserGroupInfo struct {
    Id          uint64 `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    Setting     string `json:"setting"`
    CreatedAt   int64  `json:"created_at"`
    UpdatedAt   int64  `json:"updated_at"`
}

type GetUserGroupResp = UserGroupInfo

type GetUserGroupListReq struct {
    Page     int32 `json:"page"`
    PageSize int32 `json:"page_size"`
}

type UserGroupListItem = UserGroupInfo

type GetUserGroupListResp struct {
    Total int64                 `json:"total"`
    List  []*UserGroupListItem  `json:"list"`
}
