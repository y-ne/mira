package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Maintainer struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Email    string `json:"email"`
}

func MaintainerHandler(c echo.Context) error {
	mntner := []Maintainer{
		{Name: "y-ne", Position: "Backend", Email: ""},
		{Name: "verryrw", Position: "Frontend", Email: ""},
	}

	return c.JSON(http.StatusOK, mntner)
}

func RootHandler(c echo.Context) error {
	resp := map[string]interface{}{
		"name": "Youne",
	}

	return c.Render(http.StatusOK, "index", resp)

	// return c.Render(http.StatusOK, "index", nil)
}
