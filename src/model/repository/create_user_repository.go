package repository

import (
	"context"
	"os"

	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/configuration/rest_err"
	"github.com/pultrini/CRUD-GO/src/model"
	"github.com/pultrini/CRUD-GO/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)


func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr) {
		logger.Info("Init createUser repository",
	zap.String("journey", "createUser"))

		collection_name := os.Getenv(MONGODB_USER_DB)

		collection := ur.databaseConnection.Collection(collection_name)

		value := converter.ConvertDomainToEntity(userDomain)
		
		result, err := collection.InsertOne(context.Background(), value)

		if err != nil{
			logger.Error("Error trying to create user", err, zap.String("journey", "createUser"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		value.ID = result.InsertedID.(bson.ObjectID)

		logger.Info(
			"CreateUser repository executed successfully",
			zap.String("userId", value.ID.Hex()),
			zap.String("journey", "createUser"))
		
		return converter.ConvertEntityToDomain(*value), nil
}