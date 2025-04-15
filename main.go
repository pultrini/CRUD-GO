package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pultrini/CRUD-GO/src/configuration/database/mongodb"
	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/controller"
	"github.com/pultrini/CRUD-GO/src/controller/routes"
	"github.com/pultrini/CRUD-GO/src/model/service"
)


func main(){
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()


	//Init the dependecies
	service :=service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	
	routes.InitRoutes(&router.RouterGroup, userController)
	
	if err := router.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}