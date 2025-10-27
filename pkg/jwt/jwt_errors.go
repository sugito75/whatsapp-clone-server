package jwt

import "errors"

var (
	ErrSigningMethod = errors.New("signing method is invalid")
	ErrTokenInvalid  = errors.New("token is invalid")
	ErrTokenExpired  = errors.New("token is expired")
	ErrClaimFormat   = errors.New("invalid claim format")
)
