package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenTTL  = time.Now().Add(time.Minute * 15)    //15 minutes
	refreshTokenTTL = time.Now().Add(time.Hour * 24 * 30) //1 month
)

type JwtService struct {
	accessTokenSecret  []byte
	refreshTokenSecret []byte
}

func NewService() *JwtService {
	return &JwtService{
		accessTokenSecret:  []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
		refreshTokenSecret: []byte(os.Getenv("REFRESH_TOKEN_SECRET")),
	}
}

func (s *JwtService) Generate(payload UserData, isAccessToken bool) string {
	claim := AuthClaim{
		User: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: s.getExpirationTime(isAccessToken),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := t.SignedString(s.getSecret(isAccessToken))

	return tokenString
}

func (s *JwtService) Verify(token string, isAccessToken bool) (*UserData, error) {
	secret := s.refreshTokenSecret
	if isAccessToken {
		secret = s.accessTokenSecret
	}

	t, err := s.parseToken(token, secret)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}

	if !t.Valid {
		return nil, ErrTokenInvalid
	}

	claim, ok := t.Claims.(*AuthClaim)
	if !ok {
		return nil, ErrClaimFormat
	}

	return &claim.User, nil
}

func (s *JwtService) parseToken(tokenString string, secret []byte) (*jwt.Token, error) {
	var claim AuthClaim
	return jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSigningMethod
		}

		return secret, nil
	})

}

func (s *JwtService) getSecret(isAccessToken bool) []byte {
	if isAccessToken {
		return s.accessTokenSecret
	}

	return s.refreshTokenSecret
}

func (s *JwtService) getExpirationTime(isAccessToken bool) *jwt.NumericDate {
	if isAccessToken {
		return jwt.NewNumericDate(accessTokenTTL)
	}

	return jwt.NewNumericDate(refreshTokenTTL)
}
