package user

import "mime/multipart"

type CreateUserDTO struct {
	Username       string                `form:"username" validate:"required,max=100"`
	Phone          string                `form:"phone" validate:"required,number,max=13"`
	Password       string                `form:"password" validate:"required"`
	Bio            string                `form:"bio,omitempty"`
	ProfilePicture *multipart.FileHeader `form:"profilePicture"`
}

type GetUserCredentialsDTO struct {
	Phone    string `json:"phone" validate:"required,number,max=13"`
	Password string `json:"password" validate:"required"`
}

type UserCredentialsDTO struct {
	ID             uint    `json:"id"`
	Username       string  `json:"username"`
	Phone          string  `json:"phone"`
	ProfilePicture *string `json:"profilePicture"`
	Bio            string  `json:"bio"`
	SessionID      string  `json:"sessionId"`
}
