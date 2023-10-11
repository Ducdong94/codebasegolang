package service

import "codebase.sample/model"

type UserService interface {
	GetById(uint) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
}
