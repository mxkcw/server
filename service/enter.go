package service

import (
	"server/service/manage"
	"server/service/webSite"
)

type ServiceGroup struct {
	WebSiteServiceGroup webSite.WebSiteServiceGroup
	ManageServiceGroup  manage.ManageServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
