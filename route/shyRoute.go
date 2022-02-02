package route

import (
	"ginTest/controller"
	"github.com/gin-gonic/gin"
)

func ShyRoute(engine *gin.Engine) {

	group := engine.Group("/shy")
	{
		group.GET("/test", controller.Shy{Name: "shy"}.ShyCon)
	}
}
