package controller

import "github.com/gin-gonic/gin"

type Shy struct {
	Name string
}

func (s Shy) ShyCon(ctx *gin.Context) {
	ctx.String(200, "v")
}
