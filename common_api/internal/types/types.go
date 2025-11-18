package types

type GetUserReq struct {
    Id uint64 `json:"id" form:"id"`
}

type GetUserResp struct {
	Id      uint64 `json:"id"`
	Account string `json:"account"`
}

type CreateUserReq struct {
    Account  string `json:"account"`
    Password string `json:"password"`
}

type CreateUserResp struct {
    Id      uint64 `json:"id"`
    Account string `json:"account"`
}
type UserListReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
}

type UserListResp struct {
	Total int64             `json:"total"`
	List  []UserListItem    `json:"list"`
}

type UserListItem struct {
	Id      uint64 `json:"id"`
	Account string `json:"account"`
}
