package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/mxkcw/windIneLog/windIne_log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/config"
	_ "server/docs"
	"server/middleware"
	"server/router"
)

func Routers() {
	var webSiteRouter = gin.Default()
	webSiteRouter.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		webSiteRouter.Use(gin.Logger())
	}
	// 全局路由
	groupRouterSite := router.GroupRouterAll.GroupRouterWebSite
	// 打开就能使用https了
	//Router.Use(middleware.LoadTls())
	//windIne_log.LogInfof("%s", "开始使用日志")
	// 跨域
	webSiteRouter.Use(middleware.Cors())
	windIne_log.LogInfof("%s", "开始使用跨域")
	// 配置api接口文档
	webSiteRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 不需要严重api 分组
	publicGroup := webSiteRouter.Group("api")
	{
		// 用户相关接口
		groupRouterSite.InitUserRouter(publicGroup)
		// 保存连接信息
		groupRouterSite.InitFormationRouter(publicGroup)

	}
	//开启服务
	webSiteRouter.Run(config.Config.System.HttpPort)
}
