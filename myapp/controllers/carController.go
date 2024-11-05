package controllers

import (
	"myapp/models"
	"myapp/repositories"
	"myapp/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCar(c *gin.Context) {
	var newCar models.Car
	if err := c.ShouldBindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCar.UpdatedAt = time.Now()

	err := repositories.CreateCar(newCar)
	if err != nil {
		utils.Logger.Error("Failed to insert car: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert car"})
		return
	}
	c.JSON(http.StatusCreated, newCar)
}

func GetCars(c *gin.Context) {
	cars, err := repositories.GetCars()
	if err != nil {
		utils.Logger.Error("Failed to retrieve cars: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cars"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func GetCarByName(c *gin.Context) {
	name := c.Param("name")

	car, err := repositories.GetCarByName(name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		} else {
			utils.Logger.Error("Failed to retrieve car: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve car"})
		}
		return
	}

	c.JSON(http.StatusOK, car)
}

func UpdateCar(c *gin.Context) {
	name := c.Param("name")
	var updatedCar models.Car
	if err := c.ShouldBindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repositories.UpdateCar(name, updatedCar)
	if err != nil {
		utils.Logger.Error("Failed to update car: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}
	c.JSON(http.StatusOK, updatedCar)
}

func DeleteCar(c *gin.Context) {
	name := c.Param("name")
	err := repositories.DeleteCar(name)
	if err != nil {
		utils.Logger.Error("Failed to delete car: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "car deleted"})
}
