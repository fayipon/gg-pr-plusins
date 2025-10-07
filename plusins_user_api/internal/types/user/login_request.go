package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Expire int64 `json:"expire"`
}