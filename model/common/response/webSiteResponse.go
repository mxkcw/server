package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(ctx *gin.Context, data interface{}, msg string) {
	fmt.Println("data", data)
	respModel := GetResCodeDataSuccess(
		msg,
		data,
	)
	fmt.Println("respModel:", respModel)
	ctx.SecureJSON(
		http.StatusOK,
		respModel,
	)
}

func Fail(ctx *gin.Context, data interface{}, msg string, code int64) {
	respModel := GetResCodeDataError(
		msg,
		data,
		code,
	)
	ctx.SecureJSON(
		http.StatusOK,
		respModel,
	)
}
