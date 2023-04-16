package user

import (
	"context"

	"github.com/blacheinc/gotemplate/config"
	"github.com/blacheinc/gotemplate/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Find finds user by id
func (u *User) Find() error {
	if err := database.MongoDB.Collection(config.UserCollection).FindOne(context.Background(), bson.M{"_id": u.ID}).Decode(&u); err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}
	return nil
}

// Create inserts a new user into the database
func (u *User) Create() error {
	// insert new user
	r, err := database.MongoDB.Collection(config.UserCollection).InsertOne(context.Background(), u)
	if err != nil {
		return err
	}
	u.ID = r.InsertedID.(primitive.ObjectID)
	return nil
}

// Save updates user in database
func (u *User) Save() error {
	// update user
	_, err := database.MongoDB.Collection(config.UserCollection).UpdateOne(context.Background(), bson.M{"_id": u.ID}, bson.M{"$set": u})
	if err != nil {
		return err
	}
	return nil
}
