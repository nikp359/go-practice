package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

type Server struct {
	SessionStore memstore.Store
	Engine       *gin.Engine
}

func (srv *Server) mainRouter(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Man",
		"count":   count,
	})
}

func loginRouter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"action": "login",
	})
}

func logoutRouter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"action": "logout",
	})
}

func main() {

	var srv Server

	srv.SessionStore = memstore.NewStore([]byte("secretKey"))

	router := gin.Default()

	router.Use(sessions.Sessions("access", srv.SessionStore))
	router.GET("/", srv.mainRouter)
	router.POST("/login", loginRouter)
	router.POST("/logout", logoutRouter)

	srv.Engine = router
	srv.Engine.Run("localhost:8080")
}
