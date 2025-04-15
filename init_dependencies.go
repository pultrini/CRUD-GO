package main

import (
	"github.com/pultrini/CRUD-GO/src/controller"
	"github.com/pultrini/CRUD-GO/src/model/repository"
	"github.com/pultrini/CRUD-GO/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service :=service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}