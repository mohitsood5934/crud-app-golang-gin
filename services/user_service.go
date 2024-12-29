package services

import (
	"context"
	"project/database"
	"project/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" // Import the mongo package
)

// GetCollection dynamically retrieves the MongoDB collection
func GetUserCollection() *mongo.Collection {
	return database.GetCollection("users")
}

func CreateUser(user models.User) error {
	// Dynamically get the user collection
	userCollection := GetUserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	return err
}

func GetAllUsers() ([]models.User, error) {
	// Dynamically get the user collection
	userCollection := GetUserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (models.User, error) {
	// Dynamically get the user collection
	userCollection := GetUserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	return user, err
}

func UpdateUser(id string, user models.User) error {
	// Dynamically get the user collection
	userCollection := GetUserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": user})
	return err
}

func DeleteUser(id string) error {
	// Dynamically get the user collection
	userCollection := GetUserCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := userCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
