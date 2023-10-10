package request

import (
	"codebase.sample/model"
	"github.com/labstack/echo/v4"
)

type UserRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *UserRegisterRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}
