package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/libgo/gin"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/controller"
	"github.com/zcong1993/rest-go/middleware"
	"github.com/zcong1993/rest-go/model"
	"github.com/zcong1993/rest-go/mysql"
)

func init() {
	mysql.InitDB(func(db *gorm.DB) {
		db.AutoMigrate(new(model.User))
		db.AutoMigrate(new(model.Token))
	})
}

func corsConfig() cors.Config {
	c := cors.DefaultConfig()

	c.AllowAllOrigins = true
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "HEAD", "DELETE", "PATCH", "OPTIONS"}
	c.ExposeHeaders = []string{common.HEADER_TOTAL_COUNT}
	return c
}

func createGinEngine() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(corsConfig()))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	v1 := r.Group("/v1")

	{
		v1.POST("/register", ginerr.CreateGinController(controller.Register))
		v1.POST("/login", ginerr.CreateGinController(controller.Login))
	}

	auth := v1.Group("", middleware.Auth)

	{
		auth.GET("/me", controller.Me)
	}

	return r
}

func main() {
	r := createGinEngine()
	r.Run(":8080")
}
