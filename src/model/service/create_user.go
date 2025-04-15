package service

import (

	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/configuration/rest_err"
	"github.com/pultrini/CRUD-GO/src/model"
	"go.uber.org/zap"
)



func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr){
	logger.Info("Init createUser model", zap.String("journey", "createUSer"))

	userDomain.EncryptPassword()
	
	userDomainRepository, err :=ud.userRepository.CreateUser(userDomain)
	if err != nil{
		logger.Error("Init createUser model", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service excecuted successfully",
		zap.String("userID", userDomain.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}