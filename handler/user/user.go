package user

import (
	"net/http"

	"codebase.sample/dto/request"
	"codebase.sample/dto/ressponse"
	"codebase.sample/handler"
	"codebase.sample/model"
	"codebase.sample/utils"

	"github.com/labstack/echo/v4"
)

func (h *handler.Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &request.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.UserService.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, ressponse.NewUserResponse(&u))

}
