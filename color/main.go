package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Server Main struct
type Server struct {
	Store  *ColorStore
	Engine *gin.Engine
}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{
				"message": "Colors API try /color",
			},
		)
	})
	r.GET("/color", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			s.Store.GetAll(),
		)
	})
	r.POST("/color", func(c *gin.Context) {
		var color Color
		if err := c.ShouldBindJSON(&color); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		s.Store.Set(color)

		c.JSON(
			http.StatusOK,
			color,
		)
	})

	return r
}

func main() {

	var server Server

	server.Store = NewColors()
	server.Engine = server.setupRouter()

	server.Engine.Run("localhost:8080")
}
