package webSite

import "server/service"

type WebSiteControllerGroup struct {
	UserApi
	FormationApi
}

var (
	userService      = service.ServiceGroupApp.WebSiteServiceGroup.UserService
	formationService = service.ServiceGroupApp.WebSiteServiceGroup.FormationService
)
