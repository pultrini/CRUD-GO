package service

import (
	"fmt"

	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/configuration/rest_err"
	"github.com/pultrini/CRUD-GO/src/model"
	"go.uber.org/zap"
)



func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr{
	logger.Info("Init createUser model", zap.String("journey", "createUSer"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	
	return nil
}