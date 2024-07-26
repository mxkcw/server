package webSite

import (
	"github.com/gin-gonic/gin"
	"server/controller/v1"
	"server/middleware"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("user/v1")
	userApi := v1.ControllerGroupApp.WebSiteControllerGroup.UserApi
	{
		userRouter.POST("/login", userApi.Login)
		userRouter.POST("/addRecord", userApi.AddRecord)
		userRouter.GET("/data", userApi.GroupData)
	}
	// 权限验证
	middleWareToken := userRouter.Use(middleware.JWT()) //进行token验证
	{
		middleWareToken.GET("/info", userApi.Info)
	}
	return userRouter
}
