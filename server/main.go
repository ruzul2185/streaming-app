package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ruzul2185/streaming-app/server/routes"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello, Streaming App!")
	})

	routes.SetupUnProtedtedRoutes(router)
	routes.SetupProtedtedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
