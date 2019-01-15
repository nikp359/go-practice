package main

import (
	"go-practice/car/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type App struct {
	router *gin.Engine
	store  *store.Store
}

func (a *App) routes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Cars service",
		})
	})
	r.GET("/car", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
	return r
}

func main() {
	db, err := gorm.Open("sqlite3", "cars.db")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	defer db.Close()

	st := &store.Store{
		Db: db,
	}

	st.Migrate()

	app := &App{
		store: st,
	}
	app.router = app.routes()
	app.router.Run("localhost:8080")
}
