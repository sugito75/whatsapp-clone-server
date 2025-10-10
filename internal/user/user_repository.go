package user

import "gorm.io/gorm"

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
		return 0, result.Error
	}

	return u.ID, nil
}
