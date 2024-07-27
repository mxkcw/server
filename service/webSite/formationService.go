package webSite

import (
	"github.com/google/uuid"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/mxkcw/windIneLog/windIne_orm/WindIne_orm_mysql"
	"server/model/common"
	"server/model/common/request"
	"strings"
	"time"
)

type FormationService struct {
}

func (f *FormationService) Formation(source, medium, widgetID, campaign, pageType string) error {
	// 生成apiKey
	uuidWithDashes := uuid.New()
	appKey := strings.ReplaceAll(uuidWithDashes.String(), "-", "")
	//生成URL连接
	var url string
	if pageType == "individual" {
		url = "https://www.wynpay.io/en/" + pageType + "?appKey=" + appKey + "&utm_source=" + source + "&utm_medium=" + medium + "&utm_campaign=" + campaign + "&currency_code=" + widgetID
	} else {
		url = "https://www.wynpay.io/en?appKey=" + appKey + "&utm_source=" + source + "&utm_medium=" + medium + "&utm_campaign=" + campaign + "&currency_code=" + widgetID
	}
	// 将信息插入数据库
	//err = WindIne_orm_mysql.Instance().MysqlDB.Debug().Create(&user).Model(&account).Error
	var siteFormation common.SiteFormation
	formation := common.SiteFormation{
		ApiKey:       appKey,
		CurrencyCode: widgetID,
		UTMSource:    source,
		UTMMedium:    medium,
		UTMCampaign:  campaign,
		URL:          url,
		PageType:     pageType,
		GMTCreate:    time.Now(),
		GMTModified:  time.Now(),
	}

	err := WindIne_orm_mysql.Instance().MysqlDB.Debug().Create(&formation).Model(&siteFormation).Error
	if err != nil {
		return err
	}
	return nil

}

func (f *FormationService) FormationList(pageInfo request.PageInfo) (error error, formation []common.SiteFormation, total int64) {
	windIne_log.LogInfof("接收参数： %d,%d,%s", pageInfo.Page, pageInfo.PageSize, pageInfo.Keyword)
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	var CustomerList []common.SiteFormation
	tx := WindIne_orm_mysql.Instance().MysqlDB.Debug().Model(&common.SiteFormation{}).Limit(limit).Offset(offset).Find(&CustomerList)
	if tx.Error != nil {
		return tx.Error, nil, 0
	}
	error = WindIne_orm_mysql.Instance().MysqlDB.Debug().Find(&CustomerList).Count(&total).Error
	if error != nil {
		return error, nil, 0
	}
	return nil, CustomerList, total
}
