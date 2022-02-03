package controller

import (
	"github.com/gin-gonic/gin"
	"path"
)

type Shy struct {
	Name string
}

func (s Shy) ShyCon(ctx *gin.Context) {
	ctx.String(200, "v")
}

func (s Shy) Add(ctx *gin.Context) {
	ctx.HTML(200, "updata.html", gin.H{})
}

func (s Shy) Updata(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	dst := path.Join("./static/updata", file.Filename)
	if err == nil {
		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(200, gin.H{
		"success": true,
		"dst":     dst,
	})
}
