package jwt

import "github.com/golang-jwt/jwt/v4"

// TokenInfo token info
type TokenInfo struct {
	// UserID 用户ID
	UserID int64 `json:"user_id,omitempty"`
	// Post 职位
	Post string `json:"post,omitempty"`
	// StoreID 店铺id
	StoreID int64 `json:"store_id,omitempty"`
}

// TokenClaims jwt claims
type TokenClaims struct {
	TokenInfo
	jwt.RegisteredClaims
}
