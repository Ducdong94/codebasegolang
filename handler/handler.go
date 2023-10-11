package handler

import (
	"codebase.sample/db"
	"codebase.sample/repository"
	"codebase.sample/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	UserService service.UserService
}

func NewHandler(r *echo.Echo) {
	db := db.InitConnection()
	h := &Handler{
		UserService: repository.NewUserRepository(db),
	}
	h.Register(r.Group("/v1/api"))
}
