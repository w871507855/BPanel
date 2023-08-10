package router

import "github.com/w871507855/BPanel/server/router/user"

type RouterGroup struct {
	UserRouterGroup user.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
