package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ruzul2185/streaming-app/server/controllers"
)

func SetupUnProtedtedRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())

	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
}
