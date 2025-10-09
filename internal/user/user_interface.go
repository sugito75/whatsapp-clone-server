package user

type UserRepository interface {
	CreateUser()
}

type UserService interface {
	CreateUser()
}

type UserHandler interface {
	CreateUser()
}
