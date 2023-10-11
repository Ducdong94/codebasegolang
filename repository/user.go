package repository

import (
	"context"
	"time"

	"codebase.sample/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("user"),
	}
}

func (us *UserRepository) GetById(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = us.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserRepository) GetByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	defer cancel()
	err := us.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserRepository) Create(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newUser := model.User{
		Id:       primitive.NewObjectID(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Avatar:   user.Avatar,
	}

	_, err := us.collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
