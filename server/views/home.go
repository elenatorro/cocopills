package views

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeView(router *gin.RouterGroup) {
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home.tmpl", "")
	})
}
