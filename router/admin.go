package router

import (
	"gin-skeleton/app/controller"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {

	indexController := new(controller.IndexController)
	router.GET("/", indexController.Index)
}
