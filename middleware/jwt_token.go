package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mxkcw/windIneLog/windIne_log"
	"time"
)

// jwt身份验证demo

// 设置jwt密钥secret
var jwtSecret = []byte("111111")

type Claims struct {
	ID       string `json:"Id"`
	Phone    string `json:"Phone"`
	Username string `json:"Username"`
	Nickname string `json:"Nickname"`
	jwt.StandardClaims
}

const expireTime = 24 * time.Hour

// GenerateToken 生成token的函数
func GenerateToken(id, phone, userName, nickname string) (string, error) {
	nowTime := time.Now().UTC()
	expireTime := nowTime.Add(expireTime)
	windIne_log.LogInfof("存储用户信息:%s ,%s,%s,%s", id, phone, userName, nickname)
	claims := Claims{
		id, // 自行添加的信息
		phone,
		userName,
		nickname,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 设置token过期时间
			Issuer:    "admin",           // 设置jwt签发者
		},
	}
	// 生成token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 验证token的函数
func ParseToken(token string) (*Claims, error) {
	// 对token的密钥进行验证
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// 判断token是否过期
	if tokenClaims != nil {
		claims, _ := tokenClaims.Claims.(*Claims)
		return claims, nil
	}
	return nil, err
}

func GetUserIdByToken(context *gin.Context) (err error, c *Claims) {
	windIne_log.LogInfof("%s", context.Request.Header.Get("Token"))
	parseToken, err := ParseToken(context.Request.Header.Get("Token"))
	if err != nil {
		return err, nil
	}
	return nil, parseToken
}
