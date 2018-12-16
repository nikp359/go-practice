package main

import (
	"go-practice/car/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func (a *App) routes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Cars service",
		})
	})
	r.GET("/cars", func(c *gin.Context) {
		cars := store.NewThreeCar()
		c.JSON(http.StatusOK, cars)
	})
	return r
}

func main() {
	app := &App{}
	app.router = app.routes()
	app.router.Run("localhost:8080")
}
