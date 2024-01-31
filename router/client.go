package router

import (
	"gin-skeleton/app/controller"

	"github.com/gin-gonic/gin"
)

func ClientRouter(router *gin.Engine) {
	indexController := new(controller.IndexController)

	api := router.Group("/admin")
	{
		api.GET("index", indexController.Index)
	}
}
