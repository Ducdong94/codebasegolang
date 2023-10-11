package main

import (
	_ "codebase.sample/docs"
	"codebase.sample/handler"
	"codebase.sample/router"
)

// @title Swagger codebase golang
// version 1.0
// @description Codebase sample API
// @title Codebase sample API

// @host 127.0.0.1:8080
// @BasePath /v1/api

// @schemes http https
// @produce application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Init router
	r := router.New()

	// Init handler
	handler.NewHandler(r)

	// Start server
	r.Logger.Fatal(r.Start(":8080"))
}
