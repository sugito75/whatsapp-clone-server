package user

import "mime/multipart"

type CreateUserDTO struct {
	Username       string                `form:"username" validate:"required,max=100"`
	Phone          string                `form:"phone" validate:"required,number,max=13"`
	Password       string                `form:"password" validate:"required"`
	Bio            string                `form:"bio,omitempty"`
	ProfilePicture *multipart.FileHeader `form:"profilePicture"`
}
