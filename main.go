package main

import (
	"codebase.sample/db"
	_ "codebase.sample/docs"
	"codebase.sample/router"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	db.InitConnection()

	v1 := r.Group("/v1/api")

	h := handler

	r.Logger.Fatal(r.Start(":8080"))
}
