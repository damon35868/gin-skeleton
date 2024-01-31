package main

import (
	router "gin-skeleton/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	router.Boot(app)

}
