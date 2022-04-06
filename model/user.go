package model

type User struct {
	ID   string
	Name string
}

type UserService interface {
	GetUser(id string) (User, error)
	CreateUser(userDetails User) (User, error)
}

type UserServiceFactory interface {
	GetUserService() UserService
}
