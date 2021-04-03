package main

import "github.com/gin-gonic/gin"

func main() {
	route_handler := gin.Default()

	route_handler.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route_handler.Run()
}
