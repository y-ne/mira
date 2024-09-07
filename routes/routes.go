package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/y-ne/mira/handlers"
)

func Router(e *echo.Echo) {
	e.GET("/", handlers.RootHandler)
	e.GET("/maintainer", handlers.MaintainerHandler)
}
