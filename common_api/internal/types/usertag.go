package types




type CreateUserTagReq struct {
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
}

type CreateUserTagResp struct {
    Id uint64 `json:"id"`
}

type UpdateUserTagParam struct {
    Id uint64 `path:"id"`
}

type UpdateUserTagReq struct {
    Id          uint64 `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
}

type UpdateUserTagResp struct {
    Success bool `json:"success"`
}


type DeleteUserTagParam struct {
    Id uint64 `path:"id"`
}

type DeleteUserTagReq struct {
    Id uint64 `json:"id" form:"id"`
}

type DeleteUserTagResp struct {
    Success bool `json:"success"`
}

type GetUserTagReq struct {
    Id uint64 `json:"id" form:"id"`
}

type UserTagInfo struct {
    Id          uint64 `json:"id"`
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
    CreatedAt   int64  `json:"created_at"`
    UpdatedAt   int64  `json:"updated_at"`
}


type GetUserTagParam struct {
    Id uint64 `path:"id"`
}

type GetUserTagResp = UserTagInfo

type GetUserTagListReq struct {
    Page     int32 `json:"page"`
    PageSize int32 `json:"page_size"`
}

type UserTagListItem = UserTagInfo

type GetUserTagListResp struct {
    Total int64               `json:"total"`
    List  []*UserTagListItem  `json:"list"`
}
