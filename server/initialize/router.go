package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/w871507855/BPanel/server/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	userRouterGroup := router.RouterGroupApp.UserRouterGroup
	userRouterGroup.InitUserRouter(r)
	return r
}
