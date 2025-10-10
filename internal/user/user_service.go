package user

import "golang.org/x/crypto/bcrypt"

type userService struct {
	repo UserRepository
}

func NewService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(u CreateUserDTO) (uint, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return 0, err
	}

	user := User{
		Username:       u.Username,
		Password:       string(hashedPassword),
		Phone:          u.Phone,
		Bio:            u.Bio,
		ProfilePicture: u.ProfilePicture.Filename,
	}

	return s.repo.CreateUser(user)

}
