package route

import (
	controller2 "ginTest/src/controller"
	"github.com/gin-gonic/gin"
)

func ShyRoute(engine *gin.Engine) {

	group := engine.Group("/shy")
	{
		group.GET("/test", controller2.Shy{Name: "shy"}.ShyCon)
	}
}
