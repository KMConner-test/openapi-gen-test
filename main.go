package main

import (
	"fmt"

	"github.com/KMConner-test/openapi-gen-test/openapi"
	"github.com/KMConner-test/openapi-gen-test/server"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	handler := server.TestHandler{}
	openapi.RegisterHandlers(e, handler)
	swaggerHandler := echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URL = "/spec"
	})
	e.GET("/spec", server.GetSwaggerSpec)
	e.GET("/swagger/*", swaggerHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 8080)))
}
