package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/shy", func(context *gin.Context) {
		context.String(200, "hello,Gin测试")
	})

	r.Run(":8050")
}
