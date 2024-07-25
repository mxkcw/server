package v1

import "server/controller/v1/webSite"

type ControllerGroup struct {
	WebSiteControllerGroup webSite.WebSiteControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)
