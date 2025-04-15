package service

import (
	"github.com/pultrini/CRUD-GO/src/configuration/rest_err"
	"github.com/pultrini/CRUD-GO/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr){
	return nil, nil
}