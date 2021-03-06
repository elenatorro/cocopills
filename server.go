package main

import (
	"github.com/gin-gonic/gin"
	"github.com/elenatorro/cocopills/server/views"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./server/templates/*.tmpl")
	router.Static("/assets", "./public/assets")

	viewRoutes := router.Group("/")
	views.HomeView(viewRoutes)

	router.Run()
}
