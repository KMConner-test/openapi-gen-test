package main

import (
	"fmt"

	"github.com/KMConner-test/openapi-gen-test/openapi"
	"github.com/KMConner-test/openapi-gen-test/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handler := server.TestHandler{}
	openapi.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 8080)))
}
