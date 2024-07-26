package common

import "time"

type SiteLog struct {
	ID          uint64    `gorm:"column:id;primary_key;auto_increment" json:"id"`        // id
	UniqueID    string    `gorm:"column:unique_id" json:"unique_id"`                     // api_key
	DeviceType  string    `gorm:"column:device_type" json:"device_type"`                 // device
	Region      string    `gorm:"column:region" json:"region"`                           // nation
	Referer     string    `gorm:"column:referer" json:"referer"`                         // form address
	UtmSource   string    `gorm:"column:utm_source" json:"utm_source"`                   // from source
	UtmMedium   string    `gorm:"column:utm_medium" json:"utm_medium"`                   // from medium
	UtmCampaign string    `gorm:"column:utm_campaign" json:"utm_campaign"`               // from campaign
	GmtCreate   time.Time `gorm:"column:gmt_create;type:datetime" json:"gmt_create"`     // creation time
	GmtModified time.Time `gorm:"column:gmt_modified;type:datetime" json:"gmt_modified"` // modified time
	Everyday    time.Time `gorm:"column:everyday;type:date" json:"everyday"`             // everyday
	Count       int       `gorm:"column:count" json:"count"`                             // frequency
}

func (SiteLog) TableName() string {
	return "site_log"
}
