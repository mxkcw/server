package webSite

import (
	"github.com/gin-gonic/gin"
	"github.com/mxkcw/windIneLog/windIne_log"
	"server/model/common/request"
	"server/model/common/response"
	"server/utils"
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
