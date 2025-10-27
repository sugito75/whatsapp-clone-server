package auth

type RegisterDTO struct {
	DisplayName    string `form:"displayName" json:"displayName" validate:"required,max=100"`
	Phone          string `form:"phone" json:"phone" validate:"required,number,max=13"`
	Password       string `form:"password" json:"password" validate:"required"`
	Bio            string `form:"bio,omitempty" json:"bio"`
	ProfilePicture string `form:"profilePicture" json:"profilePicture"`
}

type LoginDTO struct {
	Phone    string `json:"phone" validate:"required,number,max=13"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type UserCredentialsDTO struct {
	ID             uint    `json:"id"`
	Username       string  `json:"username"`
	Phone          string  `json:"phone"`
	ProfilePicture *string `json:"profilePicture"`
	Bio            string  `json:"bio"`
	SessionID      string  `json:"sessionId"`
	RefreshToken   string  `json:"refreshToken"`
	AccessToken    string  `json:"accessToken"`
}
