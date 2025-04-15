package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/configuration/validation"
	"github.com/pultrini/CRUD-GO/src/controller/model/request"
	"github.com/pultrini/CRUD-GO/src/model"
	"github.com/pultrini/CRUD-GO/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)


func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"),
)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, 
		userRequest.Name, userRequest.Age)
	
	domainResult, err := uc.service.CreateUser(domain)
	if err != nil{
		logger.Error("Error trying to call CreateUser service", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfuly",
	zap.String("userID", domainResult.GetEmail()),
	zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}