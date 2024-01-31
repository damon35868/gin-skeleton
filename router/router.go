package router

import (
	"github.com/gin-gonic/gin"
)

func Boot(router *gin.Engine) {
	ClientRouter(router)
	AdminRouter(router)
}
