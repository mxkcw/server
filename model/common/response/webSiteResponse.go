package response

import (
	"github.com/gin-gonic/gin"
	windIne "github.com/mxkcw/windIneLog"
	"net/http"
	"server/config"
)

func Ok(ctx *gin.Context, data interface{}, msg string) {
	respModel := GetResCodeDataSuccess(
		msg,
		data,
		config.CurrentRunMode == windIne.RunModeRelease,
	)
	ctx.SecureJSON(
		http.StatusOK,
		respModel,
	)
}

func Fail(ctx *gin.Context, data interface{}, msg string, code int64) {
	respModel := GetResCodeDataError(
		msg,
		data,
		config.CurrentRunMode == windIne.RunModeRelease,
		code,
	)
	ctx.SecureJSON(
		http.StatusOK,
		respModel,
	)
}
