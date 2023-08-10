package api

import v1 "github.com/w871507855/BPanel/server/api/v1"

type ApiGroup struct {
	v1.V1Group
}

var ApiGroupApp = new(ApiGroup)
