package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	app()
}

func app() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.WithFields(log.Fields{
			"Request":  string(reqBody),
			"Responce": string(resBody),
		}).Info("dump")
	}))
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Hello dude",
		})
	})
	e.POST("/request", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"receive": true,
		})
	})

	e.Logger.Fatal(e.Start("localhost:1323"))
}
