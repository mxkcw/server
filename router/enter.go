package router

import (
	"server/router/manage"
	"server/router/webSite"
)

type GroupRouter struct {
	GroupRouterWebSite webSite.GroupRouterWebSite
	GroupRouterManage  manage.GroupRouterManage
}

var GroupRouterAll = new(GroupRouter)
