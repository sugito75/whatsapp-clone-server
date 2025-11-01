package auth

import (
	"errors"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) SaveToken(userId uint64, token string) error {
	var t AuthToken
	result := r.db.First(&t, "user_id = $1 AND token = $2", userId, token)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	t = AuthToken{UserID: userId, Token: token}

	result = r.db.Create(&t)
	return result.Error
}

func (r *authRepository) RemoveToken(token string) error {
	result := r.db.Delete(AuthToken{}, "token = $1", token)
	return result.Error
}

func (r *authRepository) GetToken(userId uint64, token string) string {
	var t AuthToken
	result := r.db.First(&t, "user_id = $1 AND token = $2", userId, token)

	if result.Error != nil {
		return ""
	}

	return t.Token
}
