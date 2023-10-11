package main

import (
	"codebase.sample/db"
	_ "codebase.sample/docs"
	"codebase.sample/handler"
	"codebase.sample/repository"
	"codebase.sample/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	db := db.InitConnection()

	v1 := r.Group("/v1/api")

	us := repository.NewUserRepository(db)
	h := handler.NewHandler(us)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":8080"))
}
