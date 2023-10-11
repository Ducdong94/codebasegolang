package handler

import (
	"net/http"

	"codebase.sample/dto/request_dto"
	"codebase.sample/dto/ressponse_dto"
	"codebase.sample/model"
	"codebase.sample/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &request_dto.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.UserService.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, ressponse_dto.NewUserResponse(&u))
}

func (h *Handler) Login(c echo.Context) error {
	req := &request_dto.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.UserService.GetByUsername(req.User.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbiden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbiden())
	}
	return c.JSON(http.StatusOK, ressponse_dto.NewUserResponse(u))
}
