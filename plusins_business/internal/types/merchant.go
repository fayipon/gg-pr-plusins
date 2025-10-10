package types

type CreateMerchantReq struct {
    MerchantCode string `json:"merchantCode"`
    Name          string `json:"name"`
    ContactEmail  string `json:"contactEmail"`
    ContactPhone  string `json:"contactPhone"`
    Domain        string `json:"domain"`
}

type UpdateMerchantReq struct {
    ID            int64  `json:"id"`
    Name          string `json:"name"`
    ContactEmail  string `json:"contactEmail"`
    ContactPhone  string `json:"contactPhone"`
    Domain        string `json:"domain"`
    Status        string `json:"status"`
}

type MerchantResp struct {
    ID            int64  `json:"id"`
    MerchantCode  string `json:"merchantCode"`
    Name          string `json:"name"`
    ContactEmail  string `json:"contactEmail"`
    ContactPhone  string `json:"contactPhone"`
    Domain        string `json:"domain"`
    Status        string `json:"status"`
}
