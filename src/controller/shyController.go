package controller

import (
	"ginTest/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
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
	//1.获取上传的文件
	file, err := ctx.FormFile("file")
	if err == nil {
		//2.获取后缀名 判断类型是否正确
		extName := path.Ext(file.Filename)
		allowExMap := map[string]bool{
			".jpg": true,
			".png": true,
			".gif": true,
			"jpeg": true,
		}

		if _, ok := allowExMap[extName]; !ok {
			ctx.String(200, "上传的不合法")
			return
		}

		//3.创建图片保存目录

		day := model.GetDay()
		dir := "./static/updata/" + day

		err := os.MkdirAll(dir, 0666)
		if err != nil {
			ctx.String(200, "创建文件失败")
		}

		//4.生成文件名和文件保存的目录

		unix := model.GetUnix()
		fileName := strconv.FormatInt(unix, 10) + extName

		//5.执行上传
		dst := path.Join(dir, fileName)
		ctx.SaveUploadedFile(file, dst)
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func (s Shy) UserSel(ctx *gin.Context) {
	user := []model.User{}
	model.DB.Find(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func (s Shy) UserAdd(ctx *gin.Context) {
	user := model.User{
		Name: "zph",
		Pass: "123",
	}

	model.DB.Create(&user)

	ctx.String(http.StatusOK, "增加成功")
}
