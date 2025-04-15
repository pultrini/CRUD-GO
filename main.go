package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/pultrini/CRUD-GO/src/configuration/database/mongodb"
	"github.com/pultrini/CRUD-GO/src/configuration/logger"
	"github.com/pultrini/CRUD-GO/src/controller/routes"
)


func main(){
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil{
		log.Fatalf("Error trying to connect to database, erro=%s \n", err.Error())
		return
	}

	//Init the dependecies
	userController := InitDependencies(database)

	router := gin.Default()
	
	routes.InitRoutes(&router.RouterGroup, userController)
	
	if err := router.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}