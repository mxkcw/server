package webSite

import (
	"github.com/gin-gonic/gin"
	v1 "server/controller/v1"
	"server/middleware"
)

type Formation struct {
}

func (u *UserRouter) InitFormationRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	formationRouter := Router.Group("formation/v1")
	formation_Api := v1.ControllerGroupApp.WebSiteControllerGroup.FormationApi
	{
		formationRouter.POST("/addRecord", formation_Api.AddRecord)
	}
	middleWareToken := formationRouter.Use(middleware.JWT()) //进行token验证
	{
		middleWareToken.POST("/save", formation_Api.SaveFormation)
		middleWareToken.GET("/data", formation_Api.GroupData)
		middleWareToken.GET("/pageList", formation_Api.PageListFormation)
		middleWareToken.POST("/delUrl", formation_Api.DeleteUrl)
		middleWareToken.POST("/upUrlState", formation_Api.UpdateUrlSend)
	}
	return formationRouter
}
