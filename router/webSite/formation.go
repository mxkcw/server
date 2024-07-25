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
	middleWareToken := formationRouter.Use(middleware.JWT()) //进行token验证
	userApi := v1.ControllerGroupApp.WebSiteControllerGroup.FormationApi
	{
		middleWareToken.POST("/save", userApi.SaveFormation)
		middleWareToken.GET("/pageList", userApi.PageListFormation)
	}
	return formationRouter
}
