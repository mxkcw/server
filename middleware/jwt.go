package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取token
		token := context.GetHeader("Token")
		//var redisToken string
		//if token != "" {
		//	redisToken, _ = GetStr(token)
		//}

		if token == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "没有携带token",
				"data": "",
			})
			context.Abort()
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusUnauthorized,
					"msg":  "token验证失败",
					"data": "",
				})
				context.Abort()
				return
			} else if time.Now().UTC().Unix() > claims.StandardClaims.ExpiresAt {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusUnauthorized,
					"msg":  "token已过期",
					"data": "",
				})
				context.Abort()
				return
			}
		}
	}
}
