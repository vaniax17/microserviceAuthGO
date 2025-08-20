package api

import (
	"microserviceAuthGO/src/user"

	"github.com/labstack/echo/v4"
)

func UserEndpointsMapping(e *echo.Echo) {
	e.POST("/create", user.Create)
}
