package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/validator"
	"github.com/zcong1993/rest-go/common"
	"github.com/zcong1993/rest-go/controller"
	_ "github.com/zcong1993/rest-go/docs"
	"github.com/zcong1993/rest-go/middleware"
	"os"
)

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

	binding.Validator = new(validator.DefaultValidator)

	r.Use(cors.New(corsConfig()))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	if os.Getenv("GIN_MODE") != "release" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Static("/public", "./.public")
		r.Static("/docs", "./docs/swagger")
	} else {
		// production
		r.Use(secure.New(secure.DefaultConfig()))
	}

	v1 := r.Group("/v1")

	{
		v1.POST("/register", ginerr.CreateGinController(controller.Register))
		v1.POST("/login", ginerr.CreateGinController(controller.Login))
		v1.GET("/users_batch", controller.UsersBatch)
	}

	{
		ginhelper.BindRouter(v1, "/posts", &controller.PostView{}, &controller.PostView{}, ginhelper.ReadOnly...)
		v1.GET("/posts/:id/comments", ginerr.CreateGinController(controller.GetPostComments))
	}

	auth := v1.Group("", middleware.Auth)

	{
		auth.GET("/me", controller.Me)
		auth.POST("/posts", ginerr.CreateGinController(controller.CreatePost))
		auth.POST("/comments", ginerr.CreateGinController(controller.CreateComment))
	}

	return r
}

// @title Backend API
// @version 0.1
// @description This is our backend api server.

// @license.name MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /v1
func main() {
	r := createGinEngine()
	r.Run(":8080")
}
