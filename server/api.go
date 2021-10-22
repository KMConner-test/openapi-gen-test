package server

import (
	"github.com/KMConner-test/openapi-gen-test/openapi"
	"github.com/labstack/echo/v4"
)

type TestHandler struct {
}

func (t TestHandler) GetCustomersCustomerId(ctx echo.Context, customerId int) error {
	return ctx.JSON(200, openapi.User{
		CustomerId:   customerId,
		CustomerName: "mame",
	})
}
