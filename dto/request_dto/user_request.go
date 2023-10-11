package request_dto

import (
	"codebase.sample/model"
	"github.com/labstack/echo/v4"
)

// Signup dto
type UserRegisterRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
}

func (r *UserRegisterRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.Username
	u.Email = r.Email
	h, err := u.HashPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

// Login dto

type UserLoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
}

func (r *UserLoginRequest) Bind( c echo.Context) error  {
	if err:= c.Bind(r); err!= nil {
		return err
	}
	if err:= c.Validate(r); err != nil {
		return err;
	}
	return nil
}
