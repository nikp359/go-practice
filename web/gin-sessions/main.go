package main

import "github.com/gin-contrib/sessions"
import "github.com/gin-contrib/sessions/redis"
import "github.com/gin-gonic/gin"

func main() {
	//To Config
	secret := "kN7UowyYqBaRDeLthpTjHFI9N9I6w905"
	redisConn := "localhost:6396"
	sessionName := "megaplan_telphin"

	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", redisConn, "", []byte(secret))
	r.Use(sessions.Sessions(sessionName, store))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/incr", func(c *gin.Context) {
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
		c.JSON(200, gin.H{"count": count})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
