package user

import (
	"github.com/gin-gonic/gin"
	"github.com/w871507855/BPanel/server/response"
)

type UserApi struct {
}

func (*UserApi) UserDetail(c *gin.Context) {
	response.Succes(c)
}
