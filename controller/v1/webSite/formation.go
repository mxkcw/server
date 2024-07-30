package webSite

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"github.com/mxkcw/windIneLog/windIne_log"
	"server/model/common/request"
	"server/model/common/response"
	"server/utils"
	"strconv"
	"time"
)

type FormationApi struct {
}

// SaveFormation godoc
//
//	@Summary		SaveFormation
//	@Description	保存连接信息并生成连接
//	@Tags			SaveFormation 保存连接信息并生成连接
//	@Param			data	body	request.FormationMode	true	"phone,password"
//	@Accept			json
//	@Produce		json
//	@Router			/api/formation/v1/save [Post]
//	@Success		200	{object}	response.ResData
func (f *FormationApi) SaveFormation(c *gin.Context) {
	var param request.FormationMode
	err := c.ShouldBind(&param)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	err = utils.Verify(param, utils.Formation)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	windIne_log.LogInfof("接收数据: %s,%s,%s,%s,%s", param.UtmSource, param.UtmMedium, param.CurrencyCode, param.UtmCampaign, param.PageType)
	err = formationService.Formation(param.UtmSource, param.UtmMedium, param.CurrencyCode, param.UtmCampaign, param.PageType)
	if err != nil {
		response.Fail(c, "", err.Error(), 500)
		return
	}
	response.Ok(c, "", "success")
}

// PageListFormation godoc
//
//	@Summary		PageListFormation
//	@Description	获取列列表数据
//	@Tags			PageListFormation 获取列列表数据
//	@Param			data	query 	true	"page,size"
//	@Accept			json
//	@Produce		json
//	@Router			 /api/formation/v1/pageList [Get]
//	@Success		200	{object}	response.ResData
func (f *FormationApi) PageListFormation(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.Fail(c, "", err.Error(), 500)
		return
	}
	err, formation, total := formationService.FormationList(pageInfo)
	if err != nil {
		response.Fail(c, "", err.Error(), 500)
	}
	windIne_log.LogInfof("%v", formation)
	data := make(map[string]interface{})
	data["list"] = formation
	data["total"] = total
	response.Ok(c, data, "success")

}

// AddRecord add visit record
func (f *FormationApi) AddRecord(c *gin.Context) {
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
	err, state := formationService.InsertRecord(param)
	data := make(map[string]interface{})
	data["state"] = state
	response.Ok(c, data, "success")
}

func (f *FormationApi) GroupData(c *gin.Context) {
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

	err, result := formationService.GetGroupData(param)
	response.Ok(c, result, "success")
}

func (f *FormationApi) DeleteUrl(c *gin.Context) {
	windIne_log.LogInfof("%v", c)
	var param request.DeleteParams
	var err = c.ShouldBind(&param)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	err = utils.Verify(param, utils.DeleteParams)
	if err != nil {
		windIne_log.LogErrorf("%s", err.Error())
		response.Fail(c, "", err.Error(), 500)
		return
	}
	windIne_log.LogInfof("ID%d", param.Id)
	err, result := formationService.DeleteData(param)
	if result {
		response.Ok(c, result, "delete success")
	} else {
		response.Fail(c, result, "the connection has been blocked by share deletion", 500)
	}

}

func (f *FormationApi) UpdateUrlSend(c *gin.Context) {
	windIne_log.LogInfof("%v", c)
	var param request.UpdateParams
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
	err, result := userService.UpUrlState(param)
	response.Ok(c, result, "success")
}
