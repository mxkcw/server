package webSite

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"github.com/mxkcw/windIneLog/windIne_log"
	"server/middleware"
	"server/model/common/request"
	"server/model/common/response"
	"server/utils"
	"strconv"
	"time"
)

type UserApi struct{}

// Login godoc
//
//	@Summary		Login
//	@Description	帐号登陆
//	@Tags			Login 帐号登陆
//	@Param			data	body	request.UserLogin	true	"phone,password"
//	@Accept			json
//	@Produce		json
//	@Router			/api/user/v1/login [Post]
//	@Success		200	{object}	response.ResData
func (u *UserApi) Login(c *gin.Context) {
	var param request.UserLogin
	err := c.ShouldBind(&param)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	err = utils.Verify(param, utils.UserLogin)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	windIne_log.LogInfof("%s,%s", param.Phone, param.Password)
	err, token := userService.Login(param.Phone, param.Password)

	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token

	response.Ok(c, data, "success")

}

func (u *UserApi) Info(c *gin.Context) {
	windIne_log.LogInfof("%v", c)
	err, claims := middleware.GetUserIdByToken(c)
	if err != nil {
		response.Fail(c, "", err.Error(), 500)
	}
	windIne_log.LogInfof("%s,%s,%s", claims.Username, claims.Phone, claims.Nickname)
	data := make(map[string]interface{})
	data["id"] = claims.ID
	data["name"] = claims.Username
	data["roles"] = claims.Nickname
	response.Ok(c, data, "success")

}

// AddRecord add visit record
func (u *UserApi) AddRecord(c *gin.Context) {
	windIne_log.LogInfof("%v", c)
	var param request.AddVistRecord
	var err = c.ShouldBind(&param)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	windIne_log.LogInfof("%s,%s,%s,%s,%s,%s", param.UtmSource, param.UtmMedium, param.ApiKey, param.DeviceType, param.Region, param.Referer)
	err = utils.Verify(param, utils.RecordVerify)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	//获取ip
	ip := c.ClientIP()
	//获取设备信息
	ua := useragent.Parse(c.Request.UserAgent()).OS
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	param.DeviceType = utils.DeviceType(c.Request.Header.Get("User-Agent"))
	// address todo
	param.Region = ua
	param.Referer = ip
	err, state := userService.InsertRecord(param)
	data := make(map[string]interface{})
	data["state"] = state
	response.Ok(c, data, "success")
}

func (u *UserApi) GroupData(c *gin.Context) {
	windIne_log.LogInfof("%v", c)
	var param request.GetData
	var err = c.ShouldBind(&param)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	err = utils.Verify(param, utils.GetData)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	// 当前时间
	currentTime := time.Now()
	// 年
	param.DataYear = strconv.Itoa(currentTime.Year())
	currentMonth := int(currentTime.Month())
	// 月
	param.DataMonth = fmt.Sprintf("%02d", currentMonth)

	err, result := userService.GetGroupData(param)
	response.Ok(c, result, "success")
}
