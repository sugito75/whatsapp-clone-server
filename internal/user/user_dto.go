package user

type CreateUserDTO struct {
	DisplayName    string `form:"displayName" json:"displayName" validate:"required,max=100"`
	Phone          string `form:"phone" json:"phone" validate:"required,number,max=13"`
	Password       string `form:"password" json:"password" validate:"required"`
	Bio            string `form:"bio,omitempty" json:"bio"`
	ProfilePicture string `form:"profilePicture" json:"profilePicture"`
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

type GetUserInfoDTO struct {
	ID          uint    `json:"id"`
	DisplayName string  `json:"displayName"`
	Phone       string  `json:"phone"`
	Bio         string  `json:"bio"`
	Icon        *string `json:"icon"`
}
