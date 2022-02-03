package route

import (
	"ginTest/controller"
	"github.com/gin-gonic/gin"
)

func ShyRoute(engine *gin.Engine) {

	routerGroup := engine.Group("/")
	{
		routerGroup.GET("/add", controller.Shy{}.Add)
		routerGroup.POST("/updata", controller.Shy{}.Updata)
	}

	group := engine.Group("/shy")
	{
		group.GET("/test", controller.Shy{Name: "shy"}.ShyCon)
	}

	user := engine.Group("/u")
	{
		user.GET("/sel", controller.Shy{}.UserSel)
		user.GET("/add", controller.Shy{}.Add)
	}
}
