package dto

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID   int `json:"user_id"`
	TenantID int `json:"tenant_id"`
	jwt.RegisteredClaims
}
