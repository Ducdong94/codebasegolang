package service

import "codebase.sample/model"

type UserService interface {
	GetById(uint) (*model.User, error)
	Create(*model.User) error
}
