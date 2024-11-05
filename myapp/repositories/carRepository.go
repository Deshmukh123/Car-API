package repositories

import (
	"context"
	"myapp/config"
	"myapp/models"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func InitDB(client *mongo.Client) {
	collection = client.Database(config.Database).Collection(config.Collection)
}

func CreateCar(car models.Car) error {
	_, err := collection.InsertOne(context.TODO(), car)
	return err
}

func GetCars() ([]models.Car, error) {
	var cars []models.Car
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var car models.Car
		if err := cursor.Decode(&car); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func GetCarByName(name string) (models.Car, error) {
	var car models.Car
	filter := bson.M{"name": strings.ToLower(name)}
	err := collection.FindOne(context.TODO(), filter).Decode(&car)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func UpdateCar(name string, car models.Car) error {
	car.UpdatedAt = time.Now()
	_, err := collection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.M{"$set": car})
	return err
}

func DeleteCar(name string) error {
	_, err := collection.DeleteOne(context.TODO(), bson.M{"name": name})
	return err
}
