package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	// jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/auth")
	// guestUsers.POST("/login", h.Login)
	// guestUsers.POST("/signup", h.)

}
