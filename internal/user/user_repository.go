package user

import (
	"log/slog"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(u User) (uint, error) {
	result := r.db.Create(&u)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return 0, result.Error
	}

	return u.ID, nil
}

func (r *userRepository) GetUserByPhone(phone string) *User {
	var u User
	result := r.db.First(&u, "phone = $1", phone)

	if result.Error != nil {
		slog.Error(result.Error.Error())
		return nil
	}

	return &u
}
