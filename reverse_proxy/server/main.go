package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Hello dude",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
