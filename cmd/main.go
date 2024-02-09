package main

import (
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	openapifilter "github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	swagger, err := generated.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil

	e := echo.New()
	e.Use(echomiddleware.Logger())

	options := middleware.Options{
		Options: openapifilter.Options{
			AuthenticationFunc: openapifilter.NoopAuthenticationFunc,
		},
	}
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger, &options))

	var server generated.ServerInterface = newServer()
	generated.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
