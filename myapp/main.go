package main

import (
	"context"
	"log"
	"myapp/config"
	"myapp/controllers"
	"myapp/repositories"
	"myapp/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	utils.Logger.Info("Connected to MongoDB!")

	repositories.InitDB(client)

	r := gin.Default()

	r.POST("/cars", controllers.CreateCar)
	r.GET("/cars", controllers.GetCars)
	r.GET("/cars/:name", controllers.GetCarByName)
	r.PUT("/cars/:name", controllers.UpdateCar)
	r.DELETE("/cars/:name", controllers.DeleteCar)

	r.Run(":8080")
}
