package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ruzul2185/streaming-app/server/controllers"
	"github.com/ruzul2185/streaming-app/server/middleware"
)

func SetupProtedtedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
}
