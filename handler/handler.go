package handler

import "codebase.sample/service"

type Handler struct {
	UserService service.UserService
}

func NewHandler(
	us service.UserService,
) *Handler {
	return &Handler{
		UserService: us,
	}
}
