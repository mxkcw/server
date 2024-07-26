package webSite

import (
	"errors"
	"fmt"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/mxkcw/windIneLog/windIne_orm/WindIne_orm_mysql"
	"gorm.io/gorm"
	"server/middleware"
	"server/model/common"
	"server/model/common/request"
	"server/utils"
	"time"
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

func (u *UserService) InsertRecord(params request.AddVistRecord) (error, bool) {
	newParams := request.Record{
		AddVistRecord: params,
		Everyday:      time.Now(),
	}
	fmt.Println("输出............")
	fmt.Println(newParams.Everyday)
	fmt.Println(newParams.AddVistRecord)
	//先查询数据是否存在
	var siteLog common.SiteLog
	t := time.Now().Format("2006-01-02")

	state := WindIne_orm_mysql.Instance().MysqlDB.Debug().Where(
		"unique_id = ? and device_type=? and region=? and referer=? and utm_source=? and utm_medium=? and utm_campaign=? and everyday = ?",
		newParams.ApiKey, newParams.DeviceType, newParams.Region, newParams.Referer, newParams.UtmSource, newParams.UtmMedium, newParams.UtmCampaign, t).Find(&siteLog)
	if state.Error != nil && state.Error != gorm.ErrRecordNotFound {
		return state.Error, false
	}

	if state.RowsAffected > 0 {
		siteLog.Count += 1
		upState := WindIne_orm_mysql.Instance().MysqlDB.Debug().Save(&siteLog)
		if upState.Error != nil {
			return upState.Error, false
		}
		return nil, true
	} else {
		addLog := common.SiteLog{
			DeviceType:  newParams.DeviceType,
			Referer:     newParams.Referer,
			Region:      newParams.Region,
			UniqueID:    newParams.ApiKey,
			UtmCampaign: newParams.UtmCampaign,
			UtmMedium:   newParams.UtmMedium,
			UtmSource:   newParams.UtmSource,
			Everyday:    newParams.Everyday,
			GmtCreate:   time.Now(),
			GmtModified: time.Now(),
			Count:       1,
		}
		//向表中插入数据
		result := WindIne_orm_mysql.Instance().MysqlDB.Debug().Create(&addLog)
		if result.Error != nil {
			return result.Error, false
		}
		return nil, true
	}
}
