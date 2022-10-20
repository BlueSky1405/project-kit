package jwt

import "github.com/golang-jwt/jwt/v4"

// TokenInfo token info
type TokenInfo struct {
	// UserID 用户ID
	UserID uint `json:"user_id,omitempty"`
	// RoleID 角色ID
	RoleID uint `json:"role_id,omitempty"`
}

// TokenClaims jwt claims
type TokenClaims struct {
	TokenInfo
	jwt.RegisteredClaims
}
