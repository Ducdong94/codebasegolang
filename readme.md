# ![Codebase sample]

> ### Golang/Echo codebase containing CRUD, auth, advanced patterns, etc

This codebase was created to demonstrate a fully fledged fullstack application built with **Golang/Echo** including CRUD operations, authentication, routing, pagination, and more.

## Getting started

### Install Golang (go1.11+)

Please check the official golang installation guide before you start. [Official Documentation](https://golang.org/doc/install)
Also make sure you have installed a go1.11+ version.

### Environment Config

make sure your ~/.*shrc have those variable:

```bash
➜  echo $GOPATH
/Users/xesina/go
➜  echo $GOROOT
/usr/local/go/
➜  echo $PATH
...:/usr/local/go/bin:/Users/xesina/test//bin:/usr/local/go/bin
```

For more info and detailed instructions please check this guide: [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Init project

Clone this repository:

```bash
go mod init codebase.sample
```

### Install core dependencies

For create new project

```bash
go get github.com/labstack/echo/v4 &&
go get github.com/labstack/echo-contrib &&
go get gopkg.in/go-playground/validator.v9 &&
go get github.com/labstack/echo/v4/middleware &&
go get github.com/labstack/echo-contrib/jaegertracing
```

For existing project

```bash
➜ go mod download
```

Write code for validator.go

```bash
package router

import "gopkg.in/go-playground/validator.v9"

func NewValidator() *Validator {

	return &Validator{
		validator: validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
```

Write code for router.go
```bash
package router

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()
	return e
}
```

Write code for main.go

```bash
package main

import "codebase.sample/router"

func main() {
	r := router.New()

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
```

### Setting Up Swagger

Install dependencies

```bash
go get github.com/swaggo/echo-swagger &&
go install github.com/swaggo/swag/cmd/swag@latest
```

Init swagger docs

```bash
swag init
```

Modify main.go

```bash
package main

import (
	"codebase.sample/router"
	_ "codebase.sample/docs" // add for swagger
	echoSwagger "github.com/swaggo/echo-swagger" // add for swagger
)

func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler) // add for swagger

	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
```


### Setting Up Datasource (MongoDB)


Please check the official mongodb installation guide. [Official Documentation](https://www.mongodb.com/docs/manual/administration/install-community/)


Install dependencies

```bash
go get go.mongodb.org/mongo-driver/mongo &&
go get github.com/joho/godotenv
```

Create a .env file in the root directory

```bash
MONGOURI=mongodb+srv://<YOUR USERNAME HERE>:<YOUR PASSWORD HERE>@cluster0.e5akf.mongodb.net/myFirstDatabese?retryWrites=true&w=majority
```






Create env.go

```bash
package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
```

Create mongodb.go

```bash
package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"codebase.sample/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}

var client *mongo.Client
var database *mongo.Database

func InitConnection() {
	getInstance()
}

func getInstance() *mongo.Database {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating database connection now.")
		client, err := mongo.NewClient(options.Client().ApplyURI(configs.EnvMongoURI()))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			database = client.Database("codebasegolang")
			fmt.Println("Connected to MongoDB")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return database
}

func GetCollection(name string) *mongo.Collection {
	return getInstance().Collection(name)
}
```

Add this line to main.go

```bash
db.InitConnection()
```










### Run

```bash
➜ go run main.go
```

### Build

```bash
➜ go build
```

### Tests

```bash
➜ go test ./...
```