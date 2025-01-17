package webSite

import (
	"github.com/gin-gonic/gin"
	"github.com/mxkcw/windIneLog/windIne_log"
	"server/middleware"
	"server/model/common/request"
	"server/model/common/response"
	"server/utils"
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
