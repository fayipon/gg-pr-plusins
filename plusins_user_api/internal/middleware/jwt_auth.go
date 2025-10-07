package middleware

import (
	"github.com/zeromicro/go-zero/rest/token"
)

type JwtAuth struct {
	tokenMaker *token.JwtToken
}

func NewJwtAuth(secret string) *JwtAuth {
	return &JwtAuth{
		tokenMaker: token.NewJwtToken(secret),
	}
}

func (j *JwtAuth) GenerateToken(payload map[string]any) (string, error) {
	return j.tokenMaker.Encode(payload)
}