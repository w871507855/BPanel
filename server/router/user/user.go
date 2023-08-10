package user

import (
	"github.com/gin-gonic/gin"
	"github.com/w871507855/BPanel/server/api"
)

type UserRouter struct {
}

func (*UserRouter) InitUserRouter(r *gin.Engine) {
	group := r.Group("/api/v1")
	userGroup := api.ApiGroupApp.V1UserGroup
	group.GET("/user", userGroup.UserDetail)
}
