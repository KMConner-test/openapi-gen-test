package server

import (
	"github.com/KMConner-test/openapi-gen-test/openapi"
	"github.com/labstack/echo/v4"
)

func GetSwaggerSpec(ctx echo.Context) error {
	swg, err := openapi.GetSwagger()

	if err != nil {
		return err
	}
	jsonData, err := swg.MarshalJSON()
	if err != nil {
		return err
	}
	return ctx.Blob(200, "application/json", jsonData)
}
