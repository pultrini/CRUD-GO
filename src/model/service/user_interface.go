package service

import (
	"github.com/pultrini/CRUD-GO/src/configuration/rest_err"
	"github.com/pultrini/CRUD-GO/src/model"
)


func NewUserDomainService() UserDomainService{
	return &userDomainService{}
}
type userDomainService struct{
}

type UserDomainService interface{
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
