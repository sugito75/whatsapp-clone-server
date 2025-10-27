package jwt

import "github.com/golang-jwt/jwt/v5"

type AuthClaim struct {
	User UserData `json:"user"`
	jwt.RegisteredClaims
}

type UserData struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
}
