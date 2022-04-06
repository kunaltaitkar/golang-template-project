package service

import (
	"github.com/kunaltaitkar/golang-template-project/model"

	"github.com/google/uuid"
)

type UserServiceFactoryImpl struct{}

var _ model.UserServiceFactory = UserServiceFactoryImpl{}

func (UserServiceFactoryImpl) GetUserService() model.UserService {
	return userServiceImpl{}
}

type userServiceImpl struct{}

func (u userServiceImpl) CreateUser(userDetails model.User) (model.User, error) {
	userDetails.ID = uuid.New().String()
	return userDetails, nil
}

func (u userServiceImpl) GetUser(id string) (model.User, error) {
	return model.User{ID: id, Name: "kunal"}, nil
}
