package webSite

import (
	"errors"
	"fmt"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/mxkcw/windIneLog/windIne_orm/WindIne_orm_mysql"
	"server/middleware"
	"server/model/common"
	"server/utils"
)

type UserService struct {
}

func (u *UserService) Login(phone, password string) (error, string) {
	// 判断手机是否进行注册
	var userInfo common.UserInfo
	var count int64
	result := WindIne_orm_mysql.Instance().MysqlDB.Debug().Where("phone = ?", phone).Find(&userInfo).Count(&count)
	if result.Error != nil {
		return result.Error, ""
	}
	if count == 0 {
		return errors.New("account does not exist"), ""
	}
	windIne_log.LogInfof("查询数据 %v", userInfo)
	//帐号密码方式进行登陆
	if password == "" {
		return errors.New("missing parameters for account password login"), ""
	}
	if userInfo.Password != utils.MD5EncryptionGo(password) {
		return errors.New("wrong account or password"), ""
	}
	//生成 token
	token, err := middleware.GenerateToken(fmt.Sprintf("%d", userInfo.ID), userInfo.Phone, userInfo.Username, userInfo.NickName)
	if err != nil {
		return err, ""
	}
	return nil, token
}
