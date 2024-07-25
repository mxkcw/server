package main

import (
	"github.com/mxkcw/windIneLog"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/mxkcw/windIneLog/windIne_orm/WindIne_orm_mysql"
	"github.com/mxkcw/windIneLog/windIne_orm/windIne_orm_config"
	"path/filepath"
	"server/config"
	"server/initialize"
)

var (
	runMode    = "debug"
	configName = ""
	configPath = ""
	pKGMode    = "webSite"
)

//	@title						Gin-Server Swagger API接口文档
//	@version					v0.0.1
//	@description				wynpay
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						x-token
//	@BasePath					/
func main() {
	if runMode == "debug" {
		config.CurrentRunMode = windIne.RunModeDebug
		configName = "config-dev"
	} else if runMode == "test" {
		config.CurrentRunMode = windIne.RunModeTest
		configName = "config-test"
	} else if runMode == "release" {
		configName = "config-prod"
		config.CurrentRunMode = windIne.RunModeRelease
	}

	if pKGMode == "manage" {
		config.CurrentPKGMode = config.PKGModeWithManage
	} else if pKGMode == "webSite" {
		config.CurrentPKGMode = config.PKGModeWithMobile
	}

	//初始化日志配置打印-采用分片处理
	windIne.SetupWindIneBox(config.ProjectName, config.CurrentRunMode, "./logs", 9, windIne_log.WindIneLogSaveTypeDays, config.HTTPRequestTimeOut)
	windIne_log.LogInfof("========%s", config.CurrentPKGMode.String())
	//初始化读取配置
	config.InitConfig(configName, filepath.Join("/config", "webSite"))
	//输出日志
	windIne_log.LogInfof("========%s", config.Config.System.Domain)
	windIne_log.LogInfof("========%s", config.Config.System.AppEnv)
	windIne_log.LogInfof("========%s", config.Config.System.Version)
	//创建数据库连接
	WindIne_orm_mysql.Instance().OPenMysql(
		config.Config.MySql.UserName,
		config.Config.MySql.Password,
		config.Config.MySql.DbName,
		config.Config.MySql.DbHost,
		config.Config.MySql.DbPort,
		windIne_orm_config.WindIneORMTimeZoneShangHai,
		func(err error) {
			windIne_log.LogInfof("%s", "sql全局配置")
			WindIne_orm_mysql.Instance().MysqlDB.Debug().Omit("CreatedAt", "UpdatedAt", "DeletedAt")
			windIne_log.LogInfof("%s", "开启全局路由模式")
			initialize.Routers()
		},
	)

}
