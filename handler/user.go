package handler

import (
	"net/http"

	"codebase.sample/dto/request_dto"
	"codebase.sample/dto/response_dto"
	"codebase.sample/model"
	"codebase.sample/utils"

	"github.com/labstack/echo/v4"
)

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user
// @ID sign-up
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body request_dto.UserRegisterRequest true "User info for registration"
// @Success 201 {object} response_dto.userResponse
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /auth/signup [post]
func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &request_dto.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	newUser, err := h.UserService.Create(&u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, response_dto.NewUserResponse(newUser))
}

// Login godoc
// @Summary Login for existing user
// @Description Login for existing user
// @ID login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body request_dto.UserLoginRequest true "Credentials to use"
// @Success 200 {object} response_dto.userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /auth/login [post]
func (h *Handler) Login(c echo.Context) error {
	req := &request_dto.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.UserService.GetByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbiden())
	}
	if !u.CheckPassword(req.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbiden())
	}
	return c.JSON(http.StatusOK, response_dto.NewUserResponse(u))
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Get the current logged-in user
// @ID current-user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} response_dto.userResponse
// @Failure 400 {object} utils.Error
// @Failure 401 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /users [get]
func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.UserService.GetById(userIdFromToken(c))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, response_dto.NewUserResponse(u))
}

func userIdFromToken(c echo.Context) string {
	id, ok := c.Get("uid").(string)
	if !ok {
		return ""
	}
	return id
}
