package webSite

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mxkcw/windIneLog/windIne_log"
	"github.com/mxkcw/windIneLog/windIne_orm/WindIne_orm_mysql"
	"gorm.io/gorm"
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
		State:        int64(1),
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

func (f *FormationService) InsertRecord(params request.AddVistRecord) (error, bool) {
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
		"unique_id = ? and utm_source=? and everyday = ?",
		newParams.ApiKey, newParams.UtmSource, t).Find(&siteLog)
	if state.Error != nil && !errors.Is(state.Error, gorm.ErrRecordNotFound) {
		return state.Error, false
	}

	if state.RowsAffected > 0 {
		siteLog.Count += 1
		siteLog.GmtModified = time.Now()
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

func (f *FormationService) GetGroupData(params request.GetData) (error, map[string][]request.StatisticalData) {
	var result []request.GroupData
	if params.DataType == 1 { //day
		rs := WindIne_orm_mysql.Instance().MysqlDB.Debug().Table("site_log").
			Select("SUM(count) AS count,utm_source,everyday").
			Group("utm_source,everyday").Where("YEAR(everyday) = ? AND MONTH(everyday) = ?", params.DataYear, params.DataMonth).
			Scan(&result)
		if rs.Error != nil && !errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			return rs.Error, nil
		}
	} else if params.DataType == 2 { //month
		rs := WindIne_orm_mysql.Instance().MysqlDB.Debug().Table("site_log").
			Select("SUM(count) AS count,utm_source,everyday").
			Group("utm_source,MONTH(everyday)").
			Scan(&result)
		if rs.Error != nil && !errors.Is(rs.Error, gorm.ErrRecordNotFound) {
			return rs.Error, nil
		}
	}
	fmt.Printf("rs:---%+v\n", result)

	newData := make(map[string][]request.StatisticalData)
	if result != nil {
		for _, value := range result {
			newData[value.UtmSource] = append(newData[value.UtmSource], request.StatisticalData{Count: value.Count, Everyday: value.Everyday})
		}
	}
	return nil, newData
}

func (f *FormationService) DeleteData(params request.DeleteParams) (error, bool) {
	rs := WindIne_orm_mysql.Instance().MysqlDB.Debug().Table("site_formation").Where("state=1").Delete(&common.SiteFormation{ID: uint64(params.Id)})
	if rs.Error != nil && !errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		return rs.Error, false
	}
	if rs.RowsAffected == 0 {
		return nil, false
	}
	return nil, true
}

func (u *UserService) UpUrlState(params request.UpdateParams) (error, bool) {
	rs := WindIne_orm_mysql.Instance().MysqlDB.Debug().Table("site_formation").Where("id=?", params.Id).Update("state", 2)
	if rs.Error != nil && !errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		return rs.Error, false
	}
	if rs.RowsAffected == 0 {
		return nil, false
	}
	return nil, true
}
