package repository

import (
	"codebase.sample/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (us *UserRepository) GetById(id uint) (*model.User, error) {
	// var u model.User
	return &model.User{}, nil
}
func (us *UserRepository) GetByUsername(username string) (*model.User, error) {
	return &model.User{}, nil
}
func (us *UserRepository) Create(*model.User) error {
	return nil
}
