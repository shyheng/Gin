package main

import (
	"ginTest/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**")

	route.ShyRoute(r)

	r.Run(":8000")
}
