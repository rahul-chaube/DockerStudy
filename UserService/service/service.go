package service

import (
	"UserService/model"
	"fmt"

	"github.com/sirupsen/logrus"
)

type UserServiceInterface interface {
	AddUser(model.User) (interface{}, error)
	ListdUser() (model.User, error)
	DeleteUser(model.User) (interface{}, error)
}

type UserService struct {
	Log *logrus.Logger
}

func InitUserService(l *logrus.Logger) *UserService {
	return &UserService{
		Log: l,
	}

}

func (s *UserService) AddUser(user model.User) (interface{}, error) {
	s.Log.Debug("Add User service is called", user)

	model.UserData = append(model.UserData, user)
	return model.UserData, nil
}

func (s *UserService) DeleteUser(model.User) (interface{}, error) {
	fmt.Println(" DeleteUser service is called ")
	return nil, nil
}

func (s *UserService) ListUser() ([]model.User, error) {
	s.Log.Debug("List user ")
	return model.UserData, nil
}
