package main

import (
	"github.com/labstack/echo/v4"
	"github.com/y-ne/mira/routes"
)

func main() {
	e := echo.New()

	routes.Router(e)

	e.Logger.Fatal(e.Start(":8123"))
}
