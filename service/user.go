package service

import (
	"codebase.sample/model"
)

type UserService interface {
	GetById(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) (*model.User, error)
}
