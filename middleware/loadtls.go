package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/unrolled/secure"
)

func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			windIne_log.LogErrorf("err:%s", err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}
