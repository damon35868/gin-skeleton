package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (indexController *IndexController) Index(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, struct {
		Code    int
		Message string
	}{
		Code:    200,
		Message: "请求成功",
	})
}
