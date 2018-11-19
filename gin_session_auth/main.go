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
	User         *User
}

type User struct {
	name        string
	password    string
	accessToken string
}

func (srv *Server) mainRouter(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("accessToken")
	if v == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "No auth",
		})
		return
	}
	//session.Set("count", count)
	//session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message":     "Hello Man",
		"accessToken": v,
	})
}

func (srv *Server) loginRouter(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("accessToken", srv.User.accessToken)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"action":  "login",
		"session": session,
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
	srv.User = &User{
		name:        "Nikita",
		password:    "lalala",
		accessToken: "secret",
	}

	router := gin.Default()

	router.Use(sessions.Sessions("access", srv.SessionStore))
	router.GET("/", srv.mainRouter)
	router.GET("/login", srv.loginRouter)
	router.POST("/logout", logoutRouter)

	srv.Engine = router
	srv.Engine.Run("localhost:8080")
}
