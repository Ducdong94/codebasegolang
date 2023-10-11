package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Avatar   string             `json:"avatar,omitempty" validate:"required"`
}

func (u *User) HashPassword(pwd string) (string, error) {
	if len(pwd) == 0 {
		return "", errors.New("password should not be empty")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	return err == nil
}
