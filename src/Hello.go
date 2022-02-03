package main

import (
	"ginTest/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	//r.GET("/", func(context *gin.Context) {
	//	context.HTML(200, "index.html", gin.H{
	//		"msg": "消息",
	//	})
	//	//context.String(200,"shy")
	//})

	route.ShyRoute(r)

	r.Run(":8050")
}
