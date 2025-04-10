package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/configuration/validation"
	"github.com/pultrini/CRUD-GO/src/controller/model/request"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"),
)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
