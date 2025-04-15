package service

import (
	"github.com/pultrini/CRUD-GO/src/configuration/rest_err"
	"github.com/pultrini/CRUD-GO/src/model"
	"github.com/pultrini/CRUD-GO/src/model/repository"
)


func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService{
	return &userDomainService{userRepository}
}
type userDomainService struct{
	userRepository repository.UserRepository
}

type UserDomainService interface{
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
