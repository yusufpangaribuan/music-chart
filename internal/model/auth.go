package models

import (
	"github.com/dgrijalva/jwt-go"
)

// LoginReq struct
type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
	jwt.StandardClaims
	UserID uint64 `json:"UserID"`
}

// UserInfo struct
type UserInfo struct {
	ID uint64
}
