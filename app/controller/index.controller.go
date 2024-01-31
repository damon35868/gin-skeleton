package controller

import (
	"gin-skeleton/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (indexController *IndexController) Index(ctx *gin.Context) {

	service.NewIndexService().Index("Controller")

	ctx.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "请求成功",
	})
}
